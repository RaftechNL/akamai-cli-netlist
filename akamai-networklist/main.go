package main

import (
	"os"
	"sort"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

// AkamaiNetworkLists object format
type AkamaiNetworkLists struct {
	NetworkLists []struct {
		UpdateEpoch                int64    `json:"updateEpoch"`
		CreateEpoch                int64    `json:"createEpoch"`
		CreateDate                 int64    `json:"createDate"`
		UpdatedBy                  string   `json:"updatedBy"`
		UpdateDate                 int64    `json:"updateDate"`
		CreatedBy                  string   `json:"createdBy"`
		ProductionActivationStatus string   `json:"productionActivationStatus"`
		StagingActivationStatus    string   `json:"stagingActivationStatus"`
		Name                       string   `json:"name"`
		Type                       string   `json:"type"`
		UniqueID                   string   `json:"unique-id"`
		List                       []string `json:"list"`
		Links                      []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		SyncPoint  int `json:"sync-point"`
		NumEntries int `json:"numEntries"`
	} `json:"network_lists"`
}

var (
	edgeConfig                                            edgegrid.Config
	version                                               string
	configSection, configFile                             string
	listType                                              string
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
			Name:  "list",
			Usage: "List Network Lists",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "extended",
					Usage:       "Show extended lists",
					Destination: &extended,
				},
				cli.BoolFlag{
					Name:        "includeDeprecated",
					Usage:       "Show deprecated lists",
					Destination: &includeDeprecated,
				},
				cli.BoolFlag{
					Name:        "includeElements",
					Usage:       "Show elements in lists",
					Destination: &includeElements,
				},
				cli.StringFlag{
					Name:        "listType",
					Value:       "IP",
					Usage:       "Type of network list [ IP | GEO ]",
					Destination: &listType,
				},
			},
			Action: cmdlistNetLists,
			// Action: func(c *cli.Context) error {

			// 	m := fmt.Sprintf("%s?listType=IP&extended=%t&includeDeprecated=%t&includeElements=%t", URL, extended, includeDeprecated, includeElements)

			// 	fmt.Println(m)
			// 	req, _ := client.NewRequest(edgeConfig, "GET", m, nil)
			// 	resp, _ := client.Do(edgeConfig, req)

			// 	defer resp.Body.Close()
			// 	byt, _ := ioutil.ReadAll(resp.Body)

			// 	result, err := ParseAkamaiNetworkLists(string(byt))
			// 	if err != nil {
			// 		fmt.Println("error:", err)
			// 	}

			// 	jsonRes, _ := json.MarshalIndent(result.NetworkLists, "", "  ")
			// 	fmt.Printf("%+v\n", string(jsonRes))

			// 	return nil
			// },
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
