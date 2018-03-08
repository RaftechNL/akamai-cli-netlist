package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdRemoveFromnetlist(c *cli.Context) error {
	return removeFromnetlist(c)
}

func removeFromnetlist(c *cli.Context) error {
	verifyArgumentByName(c, "id")

	netLists, _, err := apiClient.NetworkLists.RemoveNetworkListItem(listID, listItem)

	if err != nil {
		return err
	}

	fmt.Println(netLists.Message)

	return nil
}
