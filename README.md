## AWS API GW Auth by AWS Signatur v4

- This is example of how to set up AWS Signature v4 auth with AWS API GW. 
- Thank you example sign code from (smartystreets)[https://github.com/smartystreets-archives/go-aws-auth]

## Set up
$ go get github.com/PayungsakCNR/go-aws-auth

## .env Set Up
```
AWS_ACCESS_KEY_ID = XXXX
AWS_SECRET_ACCESS_KEY = YYYY
```

## Example Request
```
package main

import (
	"fmt"
	awsauth "go-aws-auth"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	url := "https://myapi.example.com/services"
	client := new(http.Client)

	req, err := http.NewRequest("GET", url, nil)

	awsauth.Sign(
		req,
		"ap-southeast-1",
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}
```



