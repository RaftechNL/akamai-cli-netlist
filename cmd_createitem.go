package main

import (
	"log"
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

	if len(c.StringSlice("items")) < 1 {
		log.Fatal("Please provide items!")

	}
	itemsToAdd := strings.Split(c.StringSlice("items")[0], ",")

	editListOpts := edgegrid.NetworkListsOptionsv2{
		List: itemsToAdd,
	}

	netLists, _, err := apiClient.NetworkListsv2.AppendListNetworkList(c.String("id"), editListOpts)

	if err != nil {
		return err
	}

	common.OutputJSON(netLists)

	return nil

}
