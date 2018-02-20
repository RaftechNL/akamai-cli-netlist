package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdActivateNetListStatus(c *cli.Context) error {
	return activateNetListStatus(c)
}

func activateNetListStatus(c *cli.Context) error {
	verifyArgumentByName(c, "id")

	activationEnvironment := "staging"

	if c.Bool("prd") {
		activationEnvironment = "production"
	}

	apiURI := fmt.Sprintf("%s/%s/status?env=%s", URL, listID, activationEnvironment)

	data := dataCall(apiURI, "GET", nil)

	if output == "json" {
		fmt.Println(data)
	} else {
		result, err := ActNetListStatusAPIRespParse(data)
		errorCheck(err)

		printTableActivationStatus(result, activationEnvironment)
	}

	return nil
}
