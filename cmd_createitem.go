package main

import (
	"fmt"
	"log"
	"strings"

	common "github.com/apiheat/akamai-cli-common"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli/v2"
)

func cmdAddItemsToNetlist(c *cli.Context) error {
	return addItemsToNetlist(c)
}

func addItemsToNetlist(c *cli.Context) error {

	var itemsToAdd []string

	//TODO: fix
	// common.VerifyArgumentByName(c, "id")

	//We are not using from-file so we need to validate arguments
	if c.String("from-file") == "" {
		//TODO: fix
		// common.VerifyArgumentByName(c, "items")

		//Add out items to slice
		itemsToAdd = strings.Split(c.String("items"), ",")

	} else {
		var errReadFile error

		path := c.String("from-file")

		itemsToAdd, errReadFile = readLinesFromFile(path)
		if errReadFile != nil {
			log.Fatal(errReadFile)
		}

	}

	editListOpts := service.NetworkListsOptionsv2{
		List: itemsToAdd,
	}

	netLists, err := apiClient.AddNetworkListElement(c.String("id"), editListOpts)

	if err != nil {
		fmt.Println(err)
		return err
	}

	common.OutputJSON(netLists)

	return nil

}
