package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
)

func cmdCreateNetList(c *cli.Context) error {
	return createNetList(c)
}

func createNetList(c *cli.Context) error {
	verifyArgumentByName(c, "id")

	apiURI := URL

	newNetworkList := SingleAkamaiNetworkList{Name: listName, Type: listType}
	newNetworkList.List = []string{}
	newNetworkList.Description = listDescription

	jsonStr, _ := json.Marshal(newNetworkList)
	var jsonObj = []byte(jsonStr)

	JSONByteArr := bytes.NewReader(jsonObj)

	data := dataCall(apiURI, "POST", JSONByteArr)

	if output == "json" {
		fmt.Println(data)
	} else {
		msg, err := NetMsgAPIRespParse(data)
		errorCheck(err)

		fmt.Println(msg.Message)
	}

	return nil
}
