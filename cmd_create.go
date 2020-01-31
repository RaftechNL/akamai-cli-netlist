package main

import (
	common "github.com/apiheat/akamai-cli-common/v4"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli/v2"
)

func cmdCreateNetList(c *cli.Context) error {
	return createNetList(c)
}

func createNetList(c *cli.Context) error {
	//TODO: fix
	// common.VerifyArgumentByName(c, "name")

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
