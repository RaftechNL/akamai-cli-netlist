package main

import (
	"errors"

	common "github.com/apiheat/akamai-cli-common"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli"
)

// cmdSyncNetListID is used by cli to sync items between source and target list
func cmdSyncNetListID(c *cli.Context) error {
	return syncNetListbyID(c)
}

// syncNetListbyID synchronizes item from src list to destination list
func syncNetListbyID(c *cli.Context) error {
	common.VerifyArgumentByName(c, "id-src")
	common.VerifyArgumentByName(c, "id-dst")

	listNetListOptsv2 := service.ListNetworkListsOptionsv2{}
	listNetListOptsv2.IncludeElements = true

	netListSrc, netlistErr := apiClient.GetNetworkList(c.String("id-src"), listNetListOptsv2)
	if netlistErr != nil {
		return netlistErr
	}

	// if we have items to sync in our source list
	// we proceed
	if len(netListSrc.List) > 0 {
		// Assign the items from src list to options obj
		syncListOpts := service.NetworkListsOptionsv2{
			List: netListSrc.List,
		}

		// Append items from src list to dst list
		netListDst, err := apiClient.AddNetworkListElement(c.String("id-dst"), syncListOpts)
		if err != nil {
			return err
		}

		common.OutputJSON(netListDst)

	}

	resultErr := service.NetworkListErrorv2{}
	resultErr.Title = "Sync failed"
	resultErr.Detail = "Source list does not contain items for sync"

	return errors.New("Source list does not have ")
}

// syncNetListWithFile synchronizes item from src list to destination list
func syncNetListWithFile(c *cli.Context) error {
	return nil
}
