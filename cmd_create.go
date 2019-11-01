package main

import (
	common "github.com/apiheat/akamai-cli-common"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli"
)

func cmdCreateNetList(c *cli.Context) error {
	return createNetList(c)
}

func createNetList(c *cli.Context) error {
	common.VerifyArgumentByName(c, "name")

	newNetworkListOpst := service.NetworkListsOptionsv2{}
	newNetworkListOpst.Description = c.String("description")
	newNetworkListOpst.Name = c.String("name")
	newNetworkListOpst.Type = c.String("type")

	newList, err := apiClient.CreateNetworkList(newNetworkListOpst)

	if err != nil {
		return err
	}

	common.OutputJSON(newList)

	return nil
}
