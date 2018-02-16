package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdRemoveFromnetlist(c *cli.Context) error {
	return removeFromnetlist(c)
}

func removeFromnetlist(c *cli.Context) error {
	verifyArgumentByName(c, "id")

	apiURI := fmt.Sprintf("%s/%s/element?element=%s", URL, listID, listItem)

	dataCall(apiURI, "DELETE", nil)

	return nil
}
