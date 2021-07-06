package main

import (
	"./core"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	demoApp()
}

func demoApp() {
	//Set the AK/SK to sign and authenticate the request.
	s := core.Signer{
		Key:    "QTWAOYTTINDUT2QVKYUC",
		Secret: "MFyfvK41ba2giqM7**********KGpownRZlmVmHc",
	}
	//The following example shows how to set the request URL and parameters to query a VPC list.

	//Specify a request method, such as GET, PUT, POST, DELETE, HEAD, and PATCH.
	//Set request host.
	//Set request URI.
	//Set parameters for the request URL.
	//Add a body if you have specified the PUT or POST method. Special characters, such as the double quotation mark ("), contained in the body must be escaped.
	r, _ := http.NewRequest("GET", "https://endpoint.example.com/v1/77b6a44cba5143ab91d13ab9a8ff44fd/vpcs?limit=1", ioutil.NopCloser(bytes.NewBuffer([]byte(""))))

	//Add header parameters, for example, x-domain-id for invoking a global service and x-project-id for invoking a project-level service.
	r.Header.Add("content-type", "application/json")
	s.Sign(r)
	fmt.Println(r.Header)
	client := http.DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
