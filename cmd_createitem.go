package main

import (
	"log"
	"strings"

	common "github.com/apiheat/akamai-cli-common"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli"
)

func cmdAddItemsToNetlist(c *cli.Context) error {
	return addItemsToNetlist(c)
}

func addItemsToNetlist(c *cli.Context) error {
	common.VerifyArgumentByName(c, "id")
	common.VerifyArgumentByName(c, "items")

	if len(c.StringSlice("items")) < 1 {
		log.Fatal("Please provide items!")

	}
	itemsToAdd := strings.Split(c.StringSlice("items")[0], ",")

	editListOpts := service.NetworkListsOptionsv2{
		List: itemsToAdd,
	}

	netLists, err := apiClient.AddNetworkListElement(c.String("id"), editListOpts)

	if err != nil {
		return err
	}

	common.OutputJSON(netLists)

	return nil

}
