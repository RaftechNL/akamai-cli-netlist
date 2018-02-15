package main

import (
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

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)

	return string(byt)
}
