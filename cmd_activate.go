package main

import (
	"fmt"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	"github.com/urfave/cli"
)

func cmdActivateNetList(c *cli.Context) error {
	return activateNetList(c)
}

func activateNetList(c *cli.Context) error {
	common.VerifyArgumentByName(c, "id")

	activationEnvironment := edgegrid.Staging

	if c.Bool("prd") {
		activationEnvironment = edgegrid.Production
	}

	actNetworkListOpts.NotificationRecipients = []string{}

	netListsActivation, _, err := apiClient.NetworkLists.ActivateNetworkList(listID, activationEnvironment, actNetworkListOpts)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(netListsActivation.Status)

	return nil

}
