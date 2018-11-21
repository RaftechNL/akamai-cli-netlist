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
	common.VerifyArgumentByName(c, "searchPattern")

	listNetListOptsv2.Search = c.String("searchPattern")
	netList, _, netlistErr := apiClient.NetworkListsv2.ListNetworkLists(listNetListOptsv2)
	if netlistErr != nil {
		return netlistErr
	}

	common.OutputJSON(netList)

	return nil
}
