package main

import (
	"os"
	"sort"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

var (
	edgeConfig                                            edgegrid.Config
	version                                               string
	configSection, configFile                             string
	listType, output                                      string
	colorOn, extended, includeDeprecated, includeElements bool
)

const (
	padding = 3
	URL     = "/network-list/v1/network_lists"
)

func main() {
	_, inCLI := os.LookupEnv("AKAMAI_CLI")

	appName := "akamai-networklist"
	if inCLI {
		appName = "akamai networklist"
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

	// Sets default value for credentials configuration file
	// to be pointing to ~/.edgerc
	dir, _ := homedir.Dir()
	dir += string(os.PathSeparator) + ".edgerc"

	// no flag => false
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "section, s",
			Value:       "default",
			Usage:       "`NAME` of section to use from credentials file",
			Destination: &configSection,
			EnvVar:      "AKAMAI_EDGERC_SECTION",
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       dir,
			Usage:       "Location of the credentials `FILE`",
			Destination: &configFile,
			EnvVar:      "AKAMAI_EDGERC_CONFIGFILE",
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
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "extended",
					Usage:       "returns more verbose data such as creation date and activation status",
					Destination: &extended,
				},
				cli.BoolFlag{
					Name:        "includeDeprecated",
					Usage:       "includes network lists that have been deleted",
					Destination: &includeDeprecated,
				},
				cli.BoolFlag{
					Name:        "includeElements",
					Usage:       "includes the full list of IP or GEO elements",
					Destination: &includeElements,
				},
				cli.StringFlag{
					Name:        "listType",
					Value:       "IP",
					Usage:       "filters by the network list type [ IP | GEO ]",
					Destination: &listType,
				},
			},
			Subcommands: []cli.Command{
				{
					Name:  "all",
					Usage: "List network lists",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "raw",
							Usage: "Show raw data of SiteShield Maps",
						},
					},
					Action: cmdlistNetLists,
				},
				{
					Name:  "list",
					Usage: "List network list by `ID`",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "output",
							Value:       "raw",
							Usage:       "Output format. Supported ['json==raw' and 'apache']",
							Destination: &output,
						},
						cli.BoolFlag{
							Name:  "only-addresses",
							Usage: "Show only Map addresses.",
						},
					},
					Action: cmdlistNetList,
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(c *cli.Context) error {

		edgeConfig = config(configFile, configSection)
		return nil
	}

	app.Run(os.Args)
}
