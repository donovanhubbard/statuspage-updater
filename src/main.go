package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type Data struct {
	Component Component `json:"component"`
}

type Component struct {
	Status string `json:"status"`
}

const UP_STATUS = "operational"
const DOWN_STATUS = "major_outage"

func main() {
	lambda.Start(run_program)
}

func run_program() {

	canary_url, present := os.LookupEnv("CANARY_URL")
	if !present {
		fmt.Println("Missing mandatory environment variable: CANARY_URL")
		os.Exit(5)
	}
	token, present := os.LookupEnv("TOKEN")
	if !present {
		fmt.Println("Missing mandatory environment variable: TOKEN")
		os.Exit(5)
	}
	page_id, present := os.LookupEnv("PAGE_ID")
	if !present {
		fmt.Println("Missing mandatory environment variable: PAGE_ID")
		os.Exit(5)
	}
	component_id, present := os.LookupEnv("COMPONENT_ID")
	if !present {
		fmt.Println("Missing mandatory environment variable: COMPONENT_ID")
		os.Exit(5)
	}

	resp, err := http.Get(canary_url)
	if err != nil {
		//Failed to get the web page.
		fmt.Println("Failed to fetch canary page. Error:")
		fmt.Println(err)
		upload_result(token, page_id, component_id, DOWN_STATUS)
	} else {
		//Any kind of return code means the web server is accessible
		fmt.Println("Retreived response from canary page.")
		fmt.Println(resp.Status)
		upload_result(token, page_id, component_id, UP_STATUS)
	}
}

func upload_result(token string, page_id string, component_id string, status string) {
	body, _ := json.Marshal(Data{
		Component: Component{
			Status: status,
		},
	})

	responseBody := bytes.NewBuffer(body)

	fmt.Println("Submiting to statuspage.io.")

	update_url := "https://api.statuspage.io/v1/pages/" + page_id + "/components/" + component_id

	req, err := http.NewRequest(http.MethodPatch, update_url, responseBody)
	if err != nil {
		fmt.Println("Failed to prepare http PATCH request. Error:")
		fmt.Println(err)
		os.Exit(1)
	}
	auth_header := "OAuth " + token
	req.Header.Add("Authorization", auth_header)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Failed to PATCH to statuspage.io api. Error:")
		fmt.Println(err)
		os.Exit(2)
	}

	defer resp.Body.Close()

	fmt.Println("Response from statuspage.io api")
	fmt.Println(resp.Status)

	// Print the response body, not needed most of the time.
	// resp_body, err := ioutil.ReadAll(resp.Body)
	//
	// if err != nil {
	// 	fmt.Println("Failed to parse response from statuspage.io api. Error:")
	// 	fmt.Println(err)
	// 	os.Exit(3)
	// }
	//
	// sb := string(resp_body)
	// fmt.Println(sb)
}
