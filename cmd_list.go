package main

import (
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
	apiURI := fmt.Sprintf("%s?listType=%s&extended=%t&includeDeprecated=%t&includeElements=%t", URL, listType, extended, includeDeprecated, includeElements)

	data := dataCall(apiURI, "GET", nil)

	if output == "json" {
		fmt.Println(data)
	} else {
		result, err := NetListsAPIRespParse(data)
		errorCheck(err)

		printTableNetworkList(result)
	}

	return nil
}

func listNetList(c *cli.Context) error {
	verifyArgumentByName(c, "id")

	apiURI := fmt.Sprintf("%s/%s?listType=%s&extended=%t&includeDeprecated=%t&includeElements=%t", URL, listID, listType, extended, includeDeprecated, includeElements)

	data := dataCall(apiURI, "GET", nil)

	if output == "json" {
		fmt.Println(data)
	} else {
		result, err := NetListAPIRespParse(data)
		errorCheck(err)

		printTableSingleNetworkList(result)
	}

	return nil
}
