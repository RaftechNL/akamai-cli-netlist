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
	activationEnvironment := "staging"

	if c.Bool("prd") {
		activationEnvironment = "production"
	}

	newNetworkListActivation := ActNetworkList{SiebelTicketID: actSiebelTicketID, Comments: actComments}
	newNetworkListActivation.NotificationRecipients = []string{}

	apiURI := fmt.Sprintf("%s/%s/activate?env=%s", URL, listID, activationEnvironment)

	jsonStr, _ := json.Marshal(newNetworkListActivation)
	var jsonObj = []byte(jsonStr)
	s := fmt.Sprintf("%s", jsonStr)
	fmt.Println("----")
	fmt.Println(s)
	fmt.Println("----")

	JSONByteArr := bytes.NewReader(jsonObj)

	fmt.Println("apiURI ----->")
	fmt.Println(apiURI)

	data := dataCall(apiURI, "POST", JSONByteArr)
	fmt.Println("data ----->")
	fmt.Println(data)

	return nil
}
