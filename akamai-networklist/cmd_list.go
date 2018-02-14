package main

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
)

func cmdlistNetLists(c *cli.Context) error {
	return listNetLists(c)
}

func listNetLists(c *cli.Context) error {
	apiURI := fmt.Sprintf("%s?listType=IP&extended=%t&includeDeprecated=%t&includeElements=%t", URL, extended, includeDeprecated, includeElements)

	data := dataGet(apiURI)

	result, err := MapsAPIRespParse(data)
	errorCheck(err)

	if c.Bool("only-ids") {
		// printIDs(result.NetworkLists)
	} else {
		jsonRes, _ := json.MarshalIndent(result.NetworkLists, "", "  ")
		fmt.Printf("%+v\n", string(jsonRes))
	}

	return nil
}

// MapsAPIRespParse whatever for now
func MapsAPIRespParse(in string) (maps AkamaiNetworkLists, err error) {
	if err = json.Unmarshal([]byte(in), &maps); err != nil {
		return
	}
	return maps, err
}
