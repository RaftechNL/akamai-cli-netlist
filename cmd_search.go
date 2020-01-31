package main

import (
	common "github.com/apiheat/akamai-cli-common"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli/v2"
)

// cmdSearchNetLists is used by cli to execute search across lists based on item
func cmdSearchNetLists(c *cli.Context) error {
	return searchNetLists(c)
}

// searchNetLists execute client API call to search for lists based on item
func searchNetLists(c *cli.Context) error {
	//TODO: fix
	// common.VerifyArgumentByName(c, "searchPattern")

	listNetListOptsv2 := service.ListNetworkListsOptionsv2{}
	listNetListOptsv2.Search = c.String("searchPattern")
	if c.String("listType") == "ANY" {
		listNetListOptsv2.TypeOflist = ""
	} else {
		listNetListOptsv2.TypeOflist = c.String("listType")
	}

	netListsRes, netlistErr := apiClient.ListNetworkLists(listNetListOptsv2)
	if netlistErr != nil {
		return netlistErr
	}

	common.OutputJSON(netListsRes)

	return nil
}
