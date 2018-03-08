package main

import (
	"fmt"

	edgegrid "github.com/RafPe/go-edgegrid"
	"github.com/urfave/cli"
)

func cmdActivateNetList(c *cli.Context) error {
	return activateNetList(c)
}

func activateNetList(c *cli.Context) error {
	verifyArgumentByName(c, "id")

	activationEnvironment := edgegrid.Staging

	if c.Bool("prd") {
		activationEnvironment = edgegrid.Production
	}

	//todo: parse notification recipents
	actNetworkListOpts.NotificationRecipients = []string{}

	netListsActivation, _, err := apiClient.NetworkLists.ActivateNetworkList(listID, activationEnvironment, actNetworkListOpts)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(netListsActivation.Status)

	return nil

}
