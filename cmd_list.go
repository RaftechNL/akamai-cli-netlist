package main

import (
	"strings"

	common "github.com/apiheat/akamai-cli-common/v4"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli/v2"
)

// cmdlistNetLists is used by cli to execute action of listing all network lists
func cmdlistNetLists(c *cli.Context) error {
	return listNetLists(c)
}

// cmdlistNetListId is used by cli to execute action of listing specific network list
func cmdlistNetListID(c *cli.Context) error {
	return listNetListbyID(c)
}

// cmdlistNetListName is used by cli to execute action of listing specific network list
func cmdlistNetListName(c *cli.Context) error {
	return listNetListbyName(c)
}

// cmdlistNetListSyncPoint is used by cli to execute action of listing specific network list
// by given SyncPoint
func cmdlistNetListSyncPoint(c *cli.Context) error {
	return listNetListbySyncPoint(c)
}

// listNetLists execute client API call to get all network lists
func listNetLists(c *cli.Context) error {

	listNetListOptsv2 := service.ListNetworkListsOptionsv2{}
	listNetListOptsv2.IncludeElements = c.Bool("includeElements")
	listNetListOptsv2.Extended = c.Bool("extended")

	// Since we are listing all we do not filter results
	listNetListOptsv2.Search = ""

	if c.String("listType") == "ANY" {

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

// listNetLists execute client API call to get specific network list
func listNetListbyID(c *cli.Context) error {
	//TODO: fix
	// common.VerifyArgumentByName(c, "id")

	listNetListOptsv2 := service.ListNetworkListsOptionsv2{}
	listNetListOptsv2.IncludeElements = c.Bool("includeElements")
	listNetListOptsv2.Extended = c.Bool("extended")

	netList, netlistErr := apiClient.GetNetworkList(c.String("id"), listNetListOptsv2)
	if netlistErr != nil {
		return netlistErr
	}

	common.OutputJSON(netList)

	return nil
}

// listNetLists execute client API call to get specific network list
// it always returns only one list - for more results use search
func listNetListbyName(c *cli.Context) error {
	//TODO: fix
	// common.VerifyArgumentByName(c, "name")

	listNetListOptsv2 := service.ListNetworkListsOptionsv2{}
	listNetListOptsv2.IncludeElements = c.Bool("includeElements")
	listNetListOptsv2.Extended = c.Bool("extended")
	listNetListOptsv2.Search = c.String("name")
	listNetListOptsv2.TypeOflist = c.String("listType")

	netList, netlistErr := apiClient.ListNetworkLists(listNetListOptsv2)
	if netlistErr != nil {
		return netlistErr
	}

	// First match wins
	for _, foundNetList := range netList.NetworkLists {
		if strings.ToLower(foundNetList.Name) == strings.ToLower(c.String("name")) {
			common.OutputJSON(foundNetList)
			return nil
		}
	}

	return nil
}

// listNetListbySyncPoint execute client API call to get specific network list
// it always returns only one list - for more results use search
func listNetListbySyncPoint(c *cli.Context) error {
	//TODO: fix
	// common.VerifyArgumentByName(c, "name")

	listNetListOptsv2 := service.ListNetworkListsOptionsv2{}
	listNetListOptsv2.IncludeElements = c.Bool("includeElements")
	listNetListOptsv2.Extended = c.Bool("extended")
	listNetListOptsv2.Search = c.String("name")
	listNetListOptsv2.TypeOflist = c.String("listType")

	netList, netlistErr := apiClient.ListNetworkLists(listNetListOptsv2)
	if netlistErr != nil {
		return netlistErr
	}

	// First match wins
	for _, foundNetList := range netList.NetworkLists {
		if strings.ToLower(foundNetList.Name) == strings.ToLower(c.String("name")) {
			common.OutputJSON(foundNetList)
			return nil
		}
	}

	return nil
}
