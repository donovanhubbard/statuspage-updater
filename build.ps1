$env:GOOS = "linux"
$env:CGO_ENABLED = "0"
$env:GOARCH = "amd64"
go build -o main .\src
~\Go\Bin\build-lambda-zip.exe -output main.zip main