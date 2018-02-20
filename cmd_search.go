package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdSearchNetLists(c *cli.Context) error {
	return searchNetLists(c)
}

func searchNetLists(c *cli.Context) error {
	verifyArgumentByName(c, "item")

	apiURI := fmt.Sprintf("%s/search?expression=%s&listType=%s&extended=%t&includeDeprecated=%t", URL, listItem, listType, extended, includeDeprecated)

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
