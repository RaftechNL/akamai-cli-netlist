package main

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
)

func cmdCreateNetList(c *cli.Context) error {
	return listNetLists(c)
}

func CreateNetList(c *cli.Context) error {
	apiURI := fmt.Sprintf("%s?listType=IP&extended=%t&includeDeprecated=%t&includeElements=%t", URL, extended, includeDeprecated, includeElements)

	data := dataCall(apiURI, "POST", nil)

	result, err := NetListAPIRespParse(data)
	errorCheck(err)

	if c.Bool("only-ids") {
		// printIDs(result.NetworkLists)
	} else {
		jsonRes, _ := json.MarshalIndent(result, "", "  ")
		fmt.Printf("%+v\n", string(jsonRes))
	}

	return nil
}
