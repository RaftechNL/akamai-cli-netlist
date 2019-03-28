package main

import (
	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	"github.com/urfave/cli"
)

func cmdCreateNetList(c *cli.Context) error {
	return createNetList(c)
}

func createNetList(c *cli.Context) error {
	common.VerifyArgumentByName(c, "name")

	newNetworkListOpst := edgegrid.NetworkListsOptionsv2{}
	newNetworkListOpst.Description = c.String("description")
	newNetworkListOpst.Name = c.String("name")
	newNetworkListOpst.Type = c.String("type")

	newList, _, err := apiClient.NetworkListsv2.CreateNetworkList(newNetworkListOpst)

	if err != nil {
		return err
	}

	common.OutputJSON(newList)

	return nil
}
