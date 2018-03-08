package main

import (
	"fmt"

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
	verifyArgumentByName(c, "item")

	netLists, resp, err := apiClient.NetworkLists.SearchNetworkListItem(listItem, listNetListOpts)

	if err != nil {
		return err
	}

	if output == "json" {
		fmt.Println(resp.Body)
	} else {
		tablePrintNetworkLists(netLists)
	}

	return nil
}
