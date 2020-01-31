package main

import (
	"log"
	"strings"

	common "github.com/apiheat/akamai-cli-common"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli/v2"
)

func cmdActivateNetList(c *cli.Context) error {
	return activateNetList(c)
}

func activateNetList(c *cli.Context) error {
	//TODO: fix
	//common.VerifyArgumentByName(c, "id")

	activationEnvironment := service.Staging

	if c.Bool("prd") {
		activationEnvironment = service.Production
	}

	if len(c.StringSlice("notificationRecipients")) < 1 {
		log.Fatal("Please provide notificationRecipients!")

	}
	notificationRecipients := strings.Split(c.StringSlice("notificationRecipients")[0], ",")

	actNetworkListOpts := service.NetworkListActivationOptsv2{
		Comments:               "",
		Fast:                   c.Bool("fast"),
		NotificationRecipients: notificationRecipients,
	}

	netListsActivation, err := apiClient.ActivateNetworkList(c.String("id"), activationEnvironment, actNetworkListOpts)

	if err != nil {
		return err
	}

	common.OutputJSON(netListsActivation)

	return nil

}
