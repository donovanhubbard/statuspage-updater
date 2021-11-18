# statuspage-updater
Checks the availability of a website and uploads the results to statuspage.io.

The very narrow scope of this project is to monitor a rasberry pi at my house running an nginx server. This program will attempt to reach my web server and upload the results to statuspage.io to see if my ISP is meeting their SLAs (not that they have any with me). The instructions to setup the webserver are here. https://github.com/fuele/statuspage-chart

In this case we are only worried about internet connectivity and not website connectivity, so any response from the web server will be marked as a success.

This program is intended to run on a schedule as a lambda function.

# How to use
## Prerequisites
* A statuspage.io account
* One page
* At least one component on that page

## Environment

First you will need your token. You can find the instructions here. https://developer.statuspage.io/#section/Authentication

Next you will need the page id. You can get that on the same page where you get your token, just down under the `Page IDs` header.

Next you will need your component ID. In statuspage.io click on your component and scroll down to the bottom of the page and look under the header `Component API ID`.

Lastly you will need the url of the web address you wish to monitor.

Once you have all of this data you need to plug it into the following mandatory environment variables.

* CANARY_URL
* TOKEN
* PAGE_ID
* COMPONENT_ID

## Deploying Code

Follow these instructions for how to deploy your go code to an aws lambda function. You can use the build.ps1 script if you are on a windows machine.
https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html


# Future additions
* Automated creating and closing incidents when the site is unavailable.
* Retrieve token from amazon's secret share rather than an environment variable
