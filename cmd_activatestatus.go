package main

import (
	common "github.com/apiheat/akamai-cli-common"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli/v2"
)

func cmdActivateNetListStatus(c *cli.Context) error {
	return activateNetListStatus(c)
}

func activateNetListStatus(c *cli.Context) error {
	//TODO: fix
	// common.VerifyArgumentByName(c, "id")

	activationEnvironment := service.Staging

	if c.Bool("prd") {
		activationEnvironment = service.Production
	}

	netListsActivationStatus, err := apiClient.GetActivationStatus(c.String("id"), activationEnvironment)

	if err != nil {
		return err
	}

	common.OutputJSON(netListsActivationStatus)

	return nil
}
