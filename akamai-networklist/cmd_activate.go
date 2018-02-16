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
	dataCall(apiURI, "POST", JSONByteArr)

	return nil
}
