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
	apiURI := fmt.Sprintf("%s/%s", URL, listID)

	newNetworkList := SingleAkamaiNetworkList{}
	sliceOfItems := strings.Split(c.StringSlice("items")[0], ",")

	// Since CLI tooling does not split our slice flag we will just split it on our own
	for _, t := range sliceOfItems {
		newNetworkList.List = append(newNetworkList.List, t)
		// fmt.Println("cos" + t)
	}

	jsonStr, _ := json.Marshal(newNetworkList)

	// s := fmt.Sprintf("%s", jsonStr)
	// fmt.Println("----")
	// fmt.Println(s)
	// fmt.Println("----")
	// // s = fmt.Sprintf("%s", c.StringSlice("items"))
	// // fmt.Println(s)
	// // fmt.Println("----")

	var jsonObj = []byte(jsonStr)

	JSONByteArr := bytes.NewReader(jsonObj)

	dataCall(apiURI, "POST", JSONByteArr)

	return nil
}
