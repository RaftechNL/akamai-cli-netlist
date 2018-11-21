package main

import (
	common "github.com/apiheat/akamai-cli-common"
	"github.com/urfave/cli"
)

func cmdRemoveNetlist(c *cli.Context) error {
	return removeNetlist(c)
}

func removeNetlist(c *cli.Context) error {
	common.VerifyArgumentByName(c, "id")

	netLists, _, err := apiClient.NetworkListsv2.DeleteNetworkList(listID)

	if err != nil {
		return err
	}

	common.OutputJSON(netLists)

	return nil
}
