package main

import (
	"fmt"
	"io"
	"io/ioutil"

	client "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
)

// dataGet which is responsible for getting information from API
func dataCall(urlPath string, method string, body io.Reader) (result string) {
	req, err := client.NewRequest(edgeConfig, method, urlPath, body)
	errorCheck(err)

	resp, err := client.Do(edgeConfig, req)
	errorCheck(err)

	s := fmt.Sprintf("RESP >>> %s", resp)
	fmt.Println("----")
	fmt.Println(s)
	fmt.Println("----")

	defer resp.Body.Close()

	byt, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(byt))
	return string(byt)
}
