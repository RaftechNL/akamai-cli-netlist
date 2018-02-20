package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

func cmdAdd2netlist(c *cli.Context) error {
	return add2netlist(c)
}

func add2netlist(c *cli.Context) error {
	verifyArgumentByName(c, "id")

	apiURI := fmt.Sprintf("%s/%s", URL, listID)

	newNetworkList := SingleAkamaiNetworkList{}
	sliceOfItems := strings.Split(c.StringSlice("items")[0], ",")

	// Since CLI tooling does not split our slice flag we will just split it on our own
	for _, t := range sliceOfItems {
		newNetworkList.List = append(newNetworkList.List, t)
		// fmt.Println("cos" + t)
	}

	jsonStr, _ := json.Marshal(newNetworkList)
	var jsonObj = []byte(jsonStr)
	JSONByteArr := bytes.NewReader(jsonObj)

	data := dataCall(apiURI, "POST", JSONByteArr)

	if output == "json" {
		fmt.Println(data)
	} else {
		msg, err := NetMsgAPIRespParse(data)
		errorCheck(err)

		fmt.Println(msg.Message)
	}

	return nil
}
