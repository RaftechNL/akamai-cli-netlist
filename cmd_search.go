package main

import (
	common "github.com/apiheat/akamai-cli-common"
	"github.com/urfave/cli"
)

// cmdSearchNetLists is used by cli to execute search across lists based on item
//
// cmd_search
func cmdSearchNetLists(c *cli.Context) error {
	return searchNetLists(c)
}

// searchNetLists execute client API call to search for lists based on item
//
// cmd_search
func searchNetLists(c *cli.Context) error {
	common.VerifyArgumentByName(c, "item")

	netLists, _, err := apiClient.NetworkLists.SearchNetworkListItem(listItem, listNetListOpts)

	if err != nil {
		return err
	}

	common.OutputJSON(netLists)

	return nil
}
