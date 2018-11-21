package main

import (
	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	"github.com/urfave/cli"
)

func cmdActivateNetListStatus(c *cli.Context) error {
	return activateNetListStatus(c)
}

func activateNetListStatus(c *cli.Context) error {
	common.VerifyArgumentByName(c, "id")

	activationEnvironment := edgegrid.Staging

	if c.Bool("prd") {
		activationEnvironment = edgegrid.Production
	}

	netListsActivationStatus, _, err := apiClient.NetworkListsv2.GetNetworkListActStatus(listID, activationEnvironment)

	if err != nil {
		return err
	}

	common.OutputJSON(netListsActivationStatus)

	return nil
}
