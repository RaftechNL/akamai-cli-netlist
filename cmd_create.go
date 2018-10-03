package main

import (
	common "github.com/apiheat/akamai-cli-common"
	"github.com/urfave/cli"
)

func cmdCreateNetList(c *cli.Context) error {
	return createNetList(c)
}

func createNetList(c *cli.Context) error {
	common.VerifyArgumentByName(c, "name")

	newList, _, err := apiClient.NetworkLists.CreateNetworkList(newNetworkListOpst)

	if err != nil {
		return err
	}

	common.OutputJSON(newList)

	return nil
}
