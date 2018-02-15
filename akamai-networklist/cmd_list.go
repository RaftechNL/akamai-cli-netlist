package main

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
)

func cmdlistNetLists(c *cli.Context) error {
	return listNetLists(c)
}

func cmdlistNetList(c *cli.Context) error {
	return listNetList(c)
}

func listNetLists(c *cli.Context) error {
	apiURI := fmt.Sprintf("%s?listType=IP&extended=%t&includeDeprecated=%t&includeElements=%t", URL, extended, includeDeprecated, includeElements)

	data := dataCall(apiURI, "GET", nil)

	result, err := NetListsAPIRespParse(data)
	errorCheck(err)

	if c.Bool("only-ids") {
		// printIDs(result.NetworkLists)
	} else {
		jsonRes, _ := json.MarshalIndent(result.NetworkLists, "", "  ")
		fmt.Printf("%+v\n", string(jsonRes))
	}

	return nil
}

func listNetList(c *cli.Context) error {
	// id := setID(c)

	apiURI := fmt.Sprintf("%s/%s?listType=IP&extended=%t&includeDeprecated=%t&includeElements=%t", URL, listID, extended, includeDeprecated, includeElements)

	data := dataCall(apiURI, "GET", nil)
	fmt.Println(data)

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
