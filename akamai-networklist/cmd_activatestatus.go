package main

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
)

func cmdActivateNetListStatus(c *cli.Context) error {
	return activateNetListStatus(c)
}

func activateNetListStatus(c *cli.Context) error {
	activationEnvironment := "staging"

	if c.Bool("prd") {
		activationEnvironment = "production"
	}

	apiURI := fmt.Sprintf("%s/%s/status?env=%s", URL, listID, activationEnvironment)

	data := dataCall(apiURI, "GET", nil)

	result, err := ActNetListStatusAPIRespParse(data)
	errorCheck(err)

	if c.Bool("only-ids") {
		// printIDs(result.NetworkLists)
	} else {
		jsonRes, _ := json.MarshalIndent(result, "", "  ")
		fmt.Printf("%+v\n", string(jsonRes))
	}

	return nil
}
