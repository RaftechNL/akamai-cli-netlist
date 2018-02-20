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

	data := dataCall(apiURI, "DELETE", nil)

	if output == "json" {
		fmt.Println(data)
	} else {
		msg, err := NetMsgAPIRespParse(data)
		errorCheck(err)

		fmt.Println(msg.Message)
	}

	return nil
}
