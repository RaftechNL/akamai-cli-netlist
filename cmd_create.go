package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdCreateNetList(c *cli.Context) error {
	return createNetList(c)
}

func createNetList(c *cli.Context) error {
	verifyArgumentByName(c, "name")

	newList, _, err := apiClient.NetworkLists.CreateNetworkList(newNetworkListOpst)

	if err != nil {
		return err
	}

	fmt.Println(newList.UniqueID)

	return nil
}
