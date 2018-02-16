package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdRemoveNetlist(c *cli.Context) error {
	return removeNetlist(c)
}

func removeNetlist(c *cli.Context) error {
	apiURI := fmt.Sprintf("%s/%s", URL, listID)

	dataCall(apiURI, "DELETE", nil)

	return nil
}
