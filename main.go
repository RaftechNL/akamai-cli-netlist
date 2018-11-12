package main

import (
	"fmt"
	"os"
	"sort"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"

	"github.com/urfave/cli"
)

var (
	apiClient          *edgegrid.Client
	apiClientOpts      *edgegrid.ClientOptions
	appVer, appName    string
	listNetListOpts    edgegrid.ListNetworkListsOptions
	actNetworkListOpts edgegrid.ActivateNetworkListOptions
	newNetworkListOpst edgegrid.CreateNetworkListOptions

	listID, listName, listDescription, listItem string
	actPrd                                      string
	listOfItems                                 []string
)

func main() {
	app := common.CreateNewApp(appName, "A CLI to interact with Akamai network lists", appVer)
	app.Flags = common.CreateFlags()
	app.Before = func(c *cli.Context) error {

		apiClientOpts := &edgegrid.ClientOptions{}
		apiClientOpts.ConfigPath = c.GlobalString("config")
		apiClientOpts.ConfigSection = c.GlobalString("section")
		apiClientOpts.DebugLevel = c.GlobalString("debug")

		// NewClient: Creates new client and returns errNewExitError
		// 			  if we failed to init
		var errNewClient error
		apiClient, errNewClient = edgegrid.NewClient(nil, apiClientOpts)

		if errNewClient != nil {
			fmt.Println(errNewClient)
			os.Exit(1)
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "List network lists objects",
			Subcommands: []cli.Command{
				{
					Name:  "all",
					Usage: "List network lists",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:        "extended",
							Usage:       "returns more verbose data such as creation date and activation status",
							Destination: &listNetListOpts.Extended,
						},
						cli.BoolFlag{
							Name:        "includeDeprecated",
							Usage:       "includes network lists that have been deleted",
							Destination: &listNetListOpts.IncludeDeprecated,
						},
						cli.BoolFlag{
							Name:        "includeElements",
							Usage:       "includes the full list of IP or GEO elements",
							Destination: &listNetListOpts.IncludeElements,
						},
						cli.StringFlag{
							Name:        "listType",
							Value:       "IP",
							Usage:       "filters by the network list type [ IP | GEO ]",
							Destination: &listNetListOpts.TypeOflist,
						},
					},
					Action: cmdlistNetLists,
				},
				{
					Name:  "list",
					Usage: "List network list by `ID`",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "id",
							Usage:       "list unique-id",
							Destination: &listID,
						},
						cli.BoolFlag{
							Name:        "extended",
							Usage:       "returns more verbose data such as creation date and activation status",
							Destination: &listNetListOpts.Extended,
						},
						cli.BoolFlag{
							Name:        "includeDeprecated",
							Usage:       "includes network lists that have been deleted",
							Destination: &listNetListOpts.IncludeDeprecated,
						},
						cli.BoolFlag{
							Name:        "includeElements",
							Usage:       "includes the full list of IP or GEO elements",
							Destination: &listNetListOpts.IncludeElements,
						},
						cli.StringFlag{
							Name:        "listType",
							Value:       "IP",
							Usage:       "filters by the network list type [ IP | GEO ]",
							Destination: &listNetListOpts.TypeOflist,
						},
					},
					Action: cmdlistNetList,
				},
			},
		},
		{
			Name:  "search",
			Usage: "search by expression",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "item",
					Usage:       "item to search for",
					Destination: &listItem,
				},
				cli.BoolFlag{
					Name:        "extended",
					Usage:       "returns more verbose data such as creation date and activation status",
					Destination: &listNetListOpts.Extended,
				},
				cli.BoolFlag{
					Name:        "includeDeprecated",
					Usage:       "includes network lists that have been deleted",
					Destination: &listNetListOpts.IncludeDeprecated,
				},
				cli.StringFlag{
					Name:        "listType",
					Value:       "IP",
					Usage:       "filters by the network list type [ IP | GEO ]",
					Destination: &listNetListOpts.TypeOflist,
				},
			},
			Action: cmdSearchNetLists,
		},
		{
			Name:  "create",
			Usage: "Creates network list",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "creates new network list",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "name",
							Value:       "",
							Usage:       "name for the new list",
							Destination: &newNetworkListOpst.Name,
						},
						cli.StringFlag{
							Name:        "description",
							Value:       "created via akamai-cli-networklist",
							Usage:       "description for the new list",
							Destination: &newNetworkListOpst.Description,
						},
						cli.StringFlag{
							Name:        "type",
							Value:       "IP",
							Usage:       "defines type of list for creation (IP/GEO)",
							Destination: &newNetworkListOpst.Type,
						},
					},
					Action: cmdCreateNetList,
				},
				{
					Name:  "item",
					Usage: "creates new network list item in list with specific`ID`",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "id",
							Usage:       "list unique-id to add item to",
							Destination: &listID,
						},
						cli.StringSliceFlag{
							Name:  "items",
							Usage: "items to be included",
						},
					},
					Action: cmdAdd2netlist,
				},
			},
		},
		{
			Name:  "activate",
			Usage: "Manage network list activation",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "activates network list",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "id",
							Usage:       "list unique-id",
							Destination: &listID,
						},
						cli.StringFlag{
							Name:        "ticket-id",
							Value:       "na",
							Usage:       "ticket for this activation",
							Destination: &actNetworkListOpts.SiebelTicketID,
						},
						cli.StringFlag{
							Name:        "comments",
							Value:       "created via akamai-cli-networklist",
							Usage:       "comment for this activation",
							Destination: &actNetworkListOpts.Comments,
						},
						cli.StringSliceFlag{
							Name:  "NotificationRecipients",
							Usage: "Notification recipients to be included in activation email",
						},
						cli.BoolFlag{
							Name:  "prd",
							Usage: "activate on production",
						},
					},
					Action: cmdActivateNetList,
				},
				{
					Name:  "status",
					Usage: "status of network list activation",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "id",
							Usage:       "list unique-id",
							Destination: &listID,
						},
						cli.BoolFlag{
							Name:  "prd",
							Usage: "activate on production",
						},
					},
					Action: cmdActivateNetListStatus,
				},
			},
		},
		{
			Name:  "remove",
			Usage: "removes network list/items",
			Subcommands: []cli.Command{
				{
					Name:  "item",
					Usage: "removes item from network list items in list with specific`ID`",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "id",
							Usage:       "list unique-id to remove item from",
							Destination: &listID,
						},
						cli.StringFlag{
							Name:        "item",
							Usage:       "item to be removed from the list",
							Destination: &listItem,
						},
					},
					Action: cmdRemoveFromnetlist,
				},
				{
					Name:  "list",
					Usage: "removes network list with specific`ID`",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "id",
							Usage:       "list unique-id to remove item from",
							Destination: &listID,
						},
					},
					Action: cmdRemoveNetlist,
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Action = func(c *cli.Context) error {

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
