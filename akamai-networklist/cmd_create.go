package main

import (
	"bytes"
	"encoding/json"

	"github.com/urfave/cli"
)

func cmdCreateNetList(c *cli.Context) error {
	return createNetList(c)
}

func createNetList(c *cli.Context) error {
	apiURI := URL

	newNetworkList := SingleAkamaiNetworkList{Name: listName, Type: listType}
	newNetworkList.List = []string{}
	newNetworkList.Description = listDescription

	jsonStr, _ := json.Marshal(newNetworkList)
	var jsonObj = []byte(jsonStr)

	JSONByteArr := bytes.NewReader(jsonObj)

	dataCall(apiURI, "POST", JSONByteArr)

	return nil
}
