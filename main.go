package main

import (
	"fmt"
	"os"
	"sort"

	common "github.com/apiheat/akamai-cli-common/v4"
	edgegrid "github.com/apiheat/go-edgegrid/v6/edgegrid"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"

	"github.com/urfave/cli/v2"
)

var (
	apiClient       *service.Netlistv2
	appVer, appName string
)

func main() {
	app := common.CreateNewApp(appName, "A CLI to interact with Akamai network lists", appVer)
	app.Flags = common.CreateFlags()

	app.Before = func(c *cli.Context) error {
		var creds *edgegrid.Credentials

		if c.String("config") != common.HomeDir() {
			var err error
			creds, err = edgegrid.NewCredentials().FromFile(c.String("config")).Section(c.String("section"))

			if err != nil {
				return err
			}
		} else {
			creds = edgegrid.NewCredentials().AutoLoad(c.String("section"))
		}

		if creds == nil {
			return fmt.Errorf("Cannot load credentials")
		}

		config := edgegrid.NewConfig().
			WithCredentials(creds).
			WithLogVerbosity(c.String("debug")).
			WithAccountSwitchKey(c.String("ask"))

		if c.String("debug") == "debug" {
			config = config.WithRequestDebug(true)
		}

		// Provide struct details needed for apiClient init
		apiClient = service.New(config)

		return nil
	}

	app.Commands = []*cli.Command{
		&cli.Command{
			Name:  "get",
			Usage: "List network lists objects",
			Subcommands: []*cli.Command{
				&cli.Command{
					Name:      "all",
					Usage:     "Gets all network list in the account",
					UsageText: fmt.Sprintf("%s get all [command options]", appName),
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:  "extended",
							Usage: "returns more verbose data such as creation date and activation status",
						},
						&cli.BoolFlag{
							Name:  "includeElements",
							Usage: "includes the full list of IP or GEO elements",
						},
						&cli.StringFlag{
							Name:  "listType",
							Value: "ANY",
							Usage: "filters by the network list type [ IP | GEO | ANY ]",
						},
					},
					Action: cmdlistNetLists,
				},
				&cli.Command{
					Name:      "by-id",
					Usage:     "Gets a network list by unique-id",
					UsageText: fmt.Sprintf("%s get by-id --id UNIQUE-ID [command options]", appName),
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "id",
							Usage: "list unique-id",
						},
						&cli.BoolFlag{
							Name:  "extended",
							Usage: "returns more verbose data such as creation date and activation status",
						},
						&cli.BoolFlag{
							Name:  "includeElements",
							Usage: "includes the full list of IP or GEO elements",
						},
					},
					Action: cmdlistNetListID,
				},
				&cli.Command{
					Name:      "by-name",
					Usage:     "Gets a network list by name",
					UsageText: fmt.Sprintf("%s get by-name --name NAME [command options]", appName),
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "list name",
						},
						&cli.BoolFlag{
							Name:  "extended",
							Usage: "returns more verbose data such as creation date and activation status",
						},
						&cli.BoolFlag{
							Name:  "includeElements",
							Usage: "includes the full list of IP or GEO elements",
						},
						&cli.StringFlag{
							Name:  "listType",
							Value: "IP",
							Usage: "filters by the network list type [ IP | GEO ]",
						},
					},
					Action: cmdlistNetListName,
				},
				&cli.Command{
					Name:      "by-syncpoint",
					Usage:     "Gets a network list by specific syncPoint",
					UsageText: fmt.Sprintf("%s get by-syncpoint --syncPoint VALUE [command options]", appName),
					Flags: []cli.Flag{
						&cli.IntFlag{
							Name:  "syncpoint",
							Usage: "Specific syncPoint of the list",
						},
						&cli.BoolFlag{
							Name:  "extended",
							Usage: "returns more verbose data such as creation date and activation status",
						},
					},
					Action: cmdlistNetListSyncPoint,
				},
			},
		},
		&cli.Command{
			Name:      "search",
			Usage:     "Finds all network lists that match specific expression ( either name or network element )",
			UsageText: fmt.Sprintf("%s search --searchPattern SEARCH-ELEMENT [command options]", appName),
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "extended",
					Usage: "returns more verbose data such as creation date and activation status",
				},
				&cli.StringFlag{
					Name:  "searchPattern",
					Usage: "includes network lists that match search pattern",
				},
				&cli.StringFlag{
					Name:  "listType",
					Value: "ANY",
					Usage: "filters by the network list type [ IP | GEO | ANY ]",
				},
			},
			Action: cmdSearchNetLists,
		},
		&cli.Command{
			Name:  "sync",
			Usage: "Synchronizes items from source list into destination list ( without activation )",
			Subcommands: []*cli.Command{
				&cli.Command{
					Name:      "aka", //TODO: Name of this command might be changed *** BETA ***
					Usage:     "Synchronizes items from source list into destination list in Akamai",
					UsageText: fmt.Sprintf("%s sync-items --id-src SOURCE-LIST-ID --id-dst TARGET-LIST-ID [command options]", appName),
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "id-src",
							Usage: "Source list ID to take items from",
						},
						&cli.StringFlag{
							Name:  "id-dst",
							Usage: "Target list ID to which items should be added",
						},
						&cli.BoolFlag{
							Name:  "force",
							Usage: "Enables removal of addresses from Akamai network",
						},
					},
					Action: cmdSyncNetListID,
				},
				&cli.Command{
					Name:      "local", //TODO: Name of this command might be changed *** BETA ***
					Usage:     "Synchronizes items from local file into destination list in Akamai",
					UsageText: fmt.Sprintf("%s sync-items --from-file PATH-TO-FILE --id-dst TARGET-LIST-ID [command options]", appName),
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "from-file",
							Usage: "Source list ID to take items from",
						},
						&cli.StringFlag{
							Name:  "id-dst",
							Usage: "Target list ID to which items should be added",
						},
						&cli.BoolFlag{
							Name:  "force",
							Usage: "Enables removal of addresses from Akamai network",
						},
					},
					Action: cmdsyncNetListWithFile,
				},
			},
		},
		&cli.Command{
			Name:  "items",
			Usage: "Manages items in network lists",
			Subcommands: []*cli.Command{
				&cli.Command{
					Name:      "add",
					Usage:     "Adds network list element to provided network list",
					UsageText: fmt.Sprintf("%s items add --id UNIQUE-ID --items ITEM1,ITEM2,ITEM3", appName),
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "id",
							Usage: "list unique-id",
						},
						&cli.StringFlag{
							Name:  "items",
							Usage: "items to be included",
						},
						&cli.StringFlag{
							Name:  "from-file",
							Usage: "items to be included from file",
						},
					},
					Action: cmdAddItemsToNetlist,
				},
				&cli.Command{
					Name:      "remove",
					Usage:     "Removes network list element from provided network list",
					UsageText: fmt.Sprintf("%s items remove --id UNIQUE-ID --element ELEMENT", appName),
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "id",
							Usage: "list unique-id",
						},
						&cli.StringFlag{
							Name:  "element",
							Usage: "element to be removed",
						},
					},
					Action: cmdRemoveItemFromNetlist,
				},
			},
		},
		&cli.Command{
			Name:      "create",
			Usage:     "Creates new network list",
			UsageText: fmt.Sprintf("%s create --name NETWORK-LIST-NAME [command options]", appName),
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Value: "",
					Usage: "name for the new list",
				},
				&cli.StringFlag{
					Name:  "description",
					Value: "created via akamai-cli-networklist",
					Usage: "description for the new list",
				},
				&cli.StringFlag{
					Name:  "type",
					Value: "IP",
					Usage: "defines type of list for creation (IP/GEO)",
				},
			},
			Action: cmdCreateNetList,
		},
		&cli.Command{
			Name:  "activate",
			Usage: "Manages network list activation/status",
			Subcommands: []*cli.Command{
				&cli.Command{
					Name:      "list",
					Usage:     "Activates network list on given network",
					UsageText: fmt.Sprintf("%s activate list --id UNIQUE-ID [command options]", appName),
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "id",
							Usage: "list unique-id",
						},
						&cli.StringFlag{
							Name:  "comments",
							Value: "activated via akamai-cli",
							Usage: "comments",
						},
						&cli.StringSliceFlag{
							Name:  "notificationRecipients",
							Usage: "recipients of notification",
						},
						&cli.BoolFlag{
							Name:  "fast",
							Usage: "n/a",
						},
						&cli.BoolFlag{
							Name:  "prd",
							Usage: "activate on production",
						},
					},
					Action: cmdActivateNetList,
				},
				&cli.Command{
					Name:      "status",
					Usage:     "Displays activation status for given network list",
					UsageText: fmt.Sprintf("%s activate status --id UNIQUE-ID [command options]", appName),
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "id",
							Usage: "list unique-id",
						},
						&cli.BoolFlag{
							Name:  "prd",
							Usage: "activate on production",
						},
					},
					Action: cmdActivateNetListStatus,
				},
			},
		},
		&cli.Command{
			Name:      "delete",
			Usage:     "Deletes network list ( ** REQUIRES LIST TO BE DEACTIVATED ON BOTH NETWORKS ** )",
			UsageText: fmt.Sprintf("%s delete --id UNIQUE-ID", appName),
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "id",
					Usage: "list unique-id to remove",
				},
			},
			Action: cmdRemoveNetlist,
		},
		&cli.Command{
			Name:      "notification",
			Usage:     "Manages network list subscription notifications ( SUBSCRIBE by default ) ",
			UsageText: fmt.Sprintf("%s notification status --id UNIQUE-ID --notificationRecipients RECIPIENTS [command options]", appName),
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:  "networkListsIDs",
					Usage: "recipients of notification",
				},
				&cli.StringSliceFlag{
					Name:  "notificationRecipients",
					Usage: "recipients of notification",
				},
				&cli.BoolFlag{
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
