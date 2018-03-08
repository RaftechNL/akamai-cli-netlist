package main

import (
	"fmt"

	edgegrid "github.com/RafPe/go-edgegrid"
	"github.com/urfave/cli"
)

func cmdActivateNetListStatus(c *cli.Context) error {
	return activateNetListStatus(c)
}

func activateNetListStatus(c *cli.Context) error {
	verifyArgumentByName(c, "id")

	activationEnvironment := edgegrid.Staging

	if c.Bool("prd") {
		activationEnvironment = edgegrid.Production
	}

	netListsActivationStatus, resp, err := apiClient.NetworkLists.GetNetworkListActivationStatus(listID, activationEnvironment)

	if err != nil {
		return err
	}

	if output == "json" {
		fmt.Println(resp.Body)
	} else {
		tablePrintNetworkListActivationStatus(netListsActivationStatus, string(activationEnvironment))
	}

	return nil
}
