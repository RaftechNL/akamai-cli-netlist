package main

import (
	"fmt"
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

	var itemsToAdd []string

	common.VerifyArgumentByName(c, "id")

	if len(c.StringSlice("from-file")) < 1 {
		common.VerifyArgumentByName(c, "items")

		if len(c.StringSlice("items")) < 1 {
			log.Fatal("Please provide items!")

		}

		//Add out items to slice
		itemsToAdd = strings.Split(c.StringSlice("items")[0], ",")

	} else {
		//TODO: This is the way we access first of the params in CLI - check if we can do it cleaner
		path := c.StringSlice("from-file")[0]

		lineItems, err := readLinesFromFile(path)
		if err != nil {
			fmt.Println(err)
		}

		for _, singleIPAddress := range lineItems {
			if singleIPAddress != "" {
				itemsToAdd = append(itemsToAdd, singleIPAddress)
			}
		}
	}

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
