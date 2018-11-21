package main

import (
	"strings"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	"github.com/urfave/cli"
)

func cmdAddItemsToNetlist(c *cli.Context) error {
	return addItemsToNetlist(c)
}

func addItemsToNetlist(c *cli.Context) error {
	common.VerifyArgumentByName(c, "id")
	common.VerifyArgumentByName(c, "items")

	editListOpts := edgegrid.NetworkListsOptionsv2{
		List: strings.Split(c.StringSlice("items")[0], ","),
	}

	netLists, _, err := apiClient.NetworkListsv2.AppendListNetworkList(listID, editListOpts)

	if err != nil {
		return err
	}

	common.OutputJSON(netLists)

	return nil

}
