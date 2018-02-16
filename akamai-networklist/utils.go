package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"

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

	defer resp.Body.Close()

	byt, _ := ioutil.ReadAll(resp.Body)

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

// getArgument gets argument from our CLI
func getArgument(c *cli.Context) string {
	var id string
	if c.NArg() == 0 {
		log.Fatal("Please provide required argument")
	}

	id = c.Args().Get(0)
	return id
}

func verifyArgumentByName(c *cli.Context, argName string) {
	if c.String(argName) == "" {
		log.Fatal("Please provide required argument(s)!")
	}
}

func printJSON(str interface{}) {
	jsonRes, _ := json.MarshalIndent(str, "", "  ")
	fmt.Printf("%+v\n", string(jsonRes))
}

func printTableNetworkList(netLists AkamaiNetworkLists) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, fmt.Sprint("# ID\tName\tNumOfentries\tStaging\tProduction"))
	for _, singleList := range netLists.NetworkLists {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%v\t%s\t%s", singleList.UniqueID, singleList.Name, singleList.NumEntries, singleList.StagingActivationStatus, singleList.ProductionActivationStatus))
	}
	w.Flush()
}

func printTableSingleNetworkList(singleList AkamaiNetworkList) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, fmt.Sprint("# ID\tName\tNumOfentries\tStaging\tProduction"))
	darwinString := "%-10s  %-15s  %-15s  %-18s  %-5s  %-17s  %s\n"
	fmt.Fprintln(w, fmt.Sprintf(darwinString, singleList.UniqueID, singleList.Name, singleList.NumEntries, singleList.StagingActivationStatus, singleList.ProductionActivationStatus))
	w.Flush()
}
