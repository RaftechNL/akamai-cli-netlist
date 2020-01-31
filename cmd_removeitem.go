package main

import (
	common "github.com/apiheat/akamai-cli-common"
	"github.com/urfave/cli/v2"
)

func cmdRemoveItemFromNetlist(c *cli.Context) error {
	return removeItemFromNetlist(c)
}

func removeItemFromNetlist(c *cli.Context) error {
	//TODO: fix
	// common.VerifyArgumentByName(c, "id")
	// common.VerifyArgumentByName(c, "element")

	netLists, err := apiClient.RemoveNetworkListElement(c.String("id"), c.String("element"))

	if err != nil {
		return err
	}

	common.OutputJSON(netLists)

	return nil
}
