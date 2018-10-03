package main

import (
	common "github.com/apiheat/akamai-cli-common"
	"github.com/urfave/cli"
)

func cmdRemoveFromnetlist(c *cli.Context) error {
	return removeFromnetlist(c)
}

func removeFromnetlist(c *cli.Context) error {
	common.VerifyArgumentByName(c, "id")

	netLists, _, err := apiClient.NetworkLists.RemoveNetworkListItem(listID, listItem)

	if err != nil {
		return err
	}

	common.OutputJSON(netLists)

	return nil
}
