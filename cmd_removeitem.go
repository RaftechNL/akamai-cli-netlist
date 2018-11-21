package main

import (
	common "github.com/apiheat/akamai-cli-common"
	"github.com/urfave/cli"
)

func cmdRemoveItemFromNetlist(c *cli.Context) error {
	return removeItemFromNetlist(c)
}

func removeItemFromNetlist(c *cli.Context) error {
	common.VerifyArgumentByName(c, "id")
	common.VerifyArgumentByName(c, "element")

	netLists, _, err := apiClient.NetworkListsv2.RemoveNetworkListElement(listID, c.String("element"))

	if err != nil {
		return err
	}

	common.OutputJSON(netLists)

	return nil
}
