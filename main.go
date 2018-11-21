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
	listNetListOptsv2  edgegrid.ListNetworkListsOptionsv2
	newNetworkListOpst edgegrid.NetworkListsOptionsv2

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
							Destination: &listNetListOptsv2.Extended,
						},
						cli.BoolFlag{
							Name:        "includeElements",
							Usage:       "includes the full list of IP or GEO elements",
							Destination: &listNetListOptsv2.IncludeElements,
						},
						cli.StringFlag{
							Name:        "listType",
							Value:       "IP",
							Usage:       "filters by the network list type [ IP | GEO ]",
							Destination: &listNetListOptsv2.TypeOflist,
						},
					},
					Action: cmdlistNetLists,
				},
				{
					Name:  "by-id",
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
							Destination: &listNetListOptsv2.Extended,
						},
						cli.BoolFlag{
							Name:        "includeElements",
							Usage:       "includes the full list of IP or GEO elements",
							Destination: &listNetListOptsv2.IncludeElements,
						},
					},
					Action: cmdlistNetListID,
				},
				{
					Name:  "by-name",
					Usage: "List network lists by `name`",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "name",
							Usage:       "list name",
							Destination: &listNetListOptsv2.Search,
						},
						cli.BoolFlag{
							Name:        "extended",
							Usage:       "returns more verbose data such as creation date and activation status",
							Destination: &listNetListOptsv2.Extended,
						},
						cli.BoolFlag{
							Name:        "includeElements",
							Usage:       "includes the full list of IP or GEO elements",
							Destination: &listNetListOptsv2.IncludeElements,
						},
						cli.StringFlag{
							Name:        "listType",
							Value:       "IP",
							Usage:       "filters by the network list type [ IP | GEO ]",
							Destination: &listNetListOptsv2.TypeOflist,
						},
					},
					Action: cmdlistNetListName,
				},
			},
		},
		{
			Name:  "search",
			Usage: "search by expression",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "extended",
					Usage:       "returns more verbose data such as creation date and activation status",
					Destination: &listNetListOptsv2.Extended,
				},
				cli.StringFlag{
					Name:        "searchPattern",
					Usage:       "includes network lists that match search pattern",
					Destination: &listNetListOptsv2.Search,
				},
				cli.StringFlag{
					Name:        "listType",
					Value:       "IP",
					Usage:       "filters by the network list type [ IP | GEO ]",
					Destination: &listNetListOptsv2.TypeOflist,
				},
			},
			Action: cmdSearchNetLists,
		},
		{
			Name:  "items",
			Usage: "manages items in network lists",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "adds items to network list",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "id",
							Usage:       "list unique-id",
							Destination: &listID,
						},
						cli.StringSliceFlag{
							Name:  "items",
							Usage: "items to be included",
						},
					},
					Action: cmdAddItemsToNetlist,
				},
				{
					Name:  "remove",
					Usage: "remove item from network list",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "id",
							Usage:       "list unique-id",
							Destination: &listID,
						},
						cli.StringFlag{
							Name:  "element",
							Usage: "element to be removed",
						},
					},
					Action: cmdRemoveItemFromNetlist,
				},
			},
		},
		{
			Name:  "create",
			Usage: "Creates network list/items",
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
			Name:  "activate",
			Usage: "Manages network list activation",
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
							Name:  "comments",
							Value: "activated via akamai-cli",
							Usage: "comments",
						},
						cli.StringSliceFlag{
							Name:  "notificationRecipients",
							Usage: "recipients of notification",
						},
						cli.BoolFlag{
							Name:  "fast",
							Usage: "n/a",
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
			Name:  "delete",
			Usage: "Delete given network list",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "id",
					Usage:       "list unique-id to remove",
					Destination: &listID,
				},
			},
			Action: cmdRemoveNetlist,
		},
		{
			Name:  "notification",
			Usage: "Manages network list notifications",
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "networkListsIDs",
					Usage: "recipients of notification",
				},
				cli.StringSliceFlag{
					Name:  "notificationRecipients",
					Usage: "recipients of notification",
				},
				cli.BoolFlag{
					Name:  "unsubscribe",
					Usage: "Unsubscribe from notifications",
				},
			},
			Action: cmdNotificationManagement,
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
