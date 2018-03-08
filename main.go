package main

import (
	"log"
	"os"
	"sort"

	"github.com/RafPe/go-edgegrid"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/urfave/cli"
)

var (
	apiClient                                   *edgegrid.Client
	apiClientOpts                               edgegrid.ClientOptions
	listNetListOpts                             edgegrid.ListNetworkListsOptions
	actNetworkListOpts                          edgegrid.ActivateNetworkListOptions
	newNetworkListOpst                          edgegrid.CreateNetworkListOptions
	version, appName, output, homeDir           string
	listID, listName, listDescription, listItem string
	actPrd                                      string
	listOfItems                                 []string
	colorOn                                     bool
)

const (
	padding = 3
)

func main() {
	_, inCLI := os.LookupEnv("AKAMAI_CLI")

	// Sets default value for credentials configuration file
	// to be pointing to ~/.edgerc
	homeDir, _ = homedir.Dir()
	homeDir += string(os.PathSeparator) + ".edgerc"

	appName := "akamai-netlist"
	if inCLI {
		appName = "akamai netlist"
	}

	app := cli.NewApp()
	app.Name = appName
	app.HelpName = appName
	app.Usage = "A CLI to interact with Akamai Network lists"
	app.Version = version
	app.Copyright = ""
	app.Authors = []cli.Author{
		{
			Name: "Petr Artamonov",
		},
		{
			Name: "Rafal Pieniazek",
		},
	}

	// no flag => false
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "section, s",
			Value:       "default",
			Usage:       "`NAME` of section to use from credentials file",
			Destination: &apiClientOpts.ConfigSection,
			EnvVar:      string(edgegrid.EnvVarEdgercSection),
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       homeDir,
			Usage:       "Location of the credentials `FILE`",
			Destination: &apiClientOpts.ConfigPath,
			EnvVar:      string(edgegrid.EnvVarEdgercPath),
		},
		cli.StringFlag{
			Name:        "output",
			Value:       "table",
			Usage:       "Defines output type ( json | table ) ",
			Destination: &output,
		},
		cli.BoolFlag{
			Name:        "no-color",
			Usage:       "Disable color output",
			Destination: &colorOn,
		},
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

	app.Before = func(c *cli.Context) error {

		apiClientOpts := &edgegrid.ClientOptions{}
		apiClientOpts.ConfigPath = getEnv(string(edgegrid.EnvVarEdgercPath), homeDir)
		apiClientOpts.ConfigSection = getEnv(string(edgegrid.EnvVarEdgercSection), "default")

		// create new API client
		apiClient = edgegrid.NewClient(nil, apiClientOpts)

		return nil
	}
	app.Run(os.Args)
}

func verifyArgumentByName(c *cli.Context, argName string) {
	if c.String(argName) == "" {
		log.Fatal("Please provide required argument(s)!")
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
