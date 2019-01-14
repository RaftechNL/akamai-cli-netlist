package main

import (
	"log"
	"strings"

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

	if len(c.StringSlice("notificationRecipients")) < 1 {
		log.Fatal("Please provide notificationRecipients!")

	}
	notificationRecipients := strings.Split(c.StringSlice("notificationRecipients")[0], ",")

	actNetworkListOpts := edgegrid.NetworkListActivationOptsv2{
		Comments: "",
		Fast:     c.Bool("fast"),
		NotificationRecipients: notificationRecipients,
	}

	netListsActivation, _, err := apiClient.NetworkListsv2.ActivateNetworkList(c.String("id"), activationEnvironment, actNetworkListOpts)

	if err != nil {
		return err
	}

	common.OutputJSON(netListsActivation)

	return nil

}
