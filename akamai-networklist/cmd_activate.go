package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
)

func cmdActivateNetList(c *cli.Context) error {
	return activateNetList(c)
}

func activateNetList(c *cli.Context) error {
	verifyArgumentByName(c, "id")

	activationEnvironment := "staging"

	if c.Bool("prd") {
		activationEnvironment = "production"
	}

	newNetworkListActivation := ActNetworkList{SiebelTicketID: actSiebelTicketID, Comments: actComments}
	newNetworkListActivation.NotificationRecipients = []string{}

	apiURI := fmt.Sprintf("%s/%s/activate?env=%s", URL, listID, activationEnvironment)

	jsonStr, _ := json.Marshal(newNetworkListActivation)
	var jsonObj = []byte(jsonStr)

	JSONByteArr := bytes.NewReader(jsonObj)

	data := dataCall(apiURI, "POST", JSONByteArr)

	if output == "json" {
		fmt.Println(data)
	} else {
		_, err := NetMsgAPIRespParse(data)
		errorCheck(err)

		fmt.Println("ok")
	}

	return nil
}
