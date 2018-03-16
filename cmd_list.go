package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// cmdlistNetLists is used by cli to execute action of listing all network lists
//
// cmd_list
func cmdlistNetLists(c *cli.Context) error {
	return listNetLists(c)
}

// cmdlistNetList is used by cli to execute action of listing specific network list
//
// cmd_list
func cmdlistNetList(c *cli.Context) error {
	return listNetList(c)
}

// listNetLists execute client API call to get all network lists
//
// cmd_list
func listNetLists(c *cli.Context) error {

	listNetListOpts.TypeOflist = "IP"
	netListsIP, resp, err := apiClient.NetworkLists.ListNetworkLists(listNetListOpts)

	listNetListOpts.TypeOflist = "GEO"
	netListsGEO, resp, err := apiClient.NetworkLists.ListNetworkLists(listNetListOpts)

	netListsResult := append(*netListsIP, *netListsGEO...)

	if err != nil {
		return err
	}

	if output == "json" {
		fmt.Println(resp.Body)
	} else {
		tablePrintNetworkLists(&netListsResult)
	}

	return nil

}

// listNetLists execute client API call to get specific network list
//
// cmd_list
func listNetList(c *cli.Context) error {
	verifyArgumentByName(c, "id")

	netList, resp, err := apiClient.NetworkLists.GetNetworkList(listID, listNetListOpts)

	if err != nil {
		return err
	}

	if output == "json" {
		fmt.Println(resp.Body)
	} else {
		tablePrintNetworkList(netList)
	}

	return nil
}
