package main

import (
	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	"github.com/urfave/cli"
)

// cmdSearchNetLists is used by cli to execute search across lists based on item
func cmdSearchNetLists(c *cli.Context) error {
	return searchNetLists(c)
}

// searchNetLists execute client API call to search for lists based on item
func searchNetLists(c *cli.Context) error {
	common.VerifyArgumentByName(c, "searchPattern")

	listNetListOptsv2 := edgegrid.ListNetworkListsOptionsv2{}
	listNetListOptsv2.Search = c.String("searchPattern")

	netList, _, netlistErr := apiClient.NetworkListsv2.ListNetworkLists(listNetListOptsv2)
	if netlistErr != nil {
		return netlistErr
	}

	common.OutputJSON(netList)

	return nil
}
