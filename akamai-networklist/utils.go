package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	client "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func errorCheck(e error) {
	if e != nil {
		color.Set(color.FgRed)
		log.Fatal(e)
		color.Unset()
	}
}

// dataGet which is responsible for getting information from API
func dataCall(urlPath string, method string, body io.Reader) (result string) {
	req, err := client.NewRequest(edgeConfig, method, urlPath, body)
	errorCheck(err)

	resp, err := client.Do(edgeConfig, req)
	errorCheck(err)

	s := fmt.Sprintf("RESP >>> %s", resp)
	fmt.Println("----")
	fmt.Println(s)
	fmt.Println("----")

	defer resp.Body.Close()

	byt, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(byt))
	return string(byt)
}

// ActNetListStatusAPIRespParse parse status of activation
func ActNetListStatusAPIRespParse(in string) (maps ActNetworkListStatus, err error) {
	if err = json.Unmarshal([]byte(in), &maps); err != nil {
		return
	}
	return maps, err
}

// NetListsAPIRespParse
func NetListsAPIRespParse(in string) (maps AkamaiNetworkLists, err error) {
	if err = json.Unmarshal([]byte(in), &maps); err != nil {
		return
	}
	return maps, err
}

// NetListAPIRespParse
func NetListAPIRespParse(in string) (maps AkamaiNetworkList, err error) {
	if err = json.Unmarshal([]byte(in), &maps); err != nil {
		return
	}
	return maps, err
}

func setID(c *cli.Context) string {
	var id string
	if c.NArg() == 0 {
		log.Fatal("Please provide ID for map")
	}

	id = c.Args().Get(0)
	return id
}

func verifyID(id string) {
	if _, err := strconv.Atoi(id); err != nil {
		errStr := fmt.Sprintf("SiteShield Map ID should be number, you provided: %q\n", id)
		log.Fatal(errStr)
	}
}

func printJSON(str interface{}) {
	jsonRes, _ := json.MarshalIndent(str, "", "  ")
	fmt.Printf("%+v\n", string(jsonRes))
}
