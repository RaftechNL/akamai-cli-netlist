package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	edgegrid "github.com/RafPe/go-edgegrid"
	"github.com/fatih/color"
)

//
// cmd_list
func tablePrintNetworkLists(netLists *[]edgegrid.AkamaiNetworkList) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)

	fmt.Fprintln(w, fmt.Sprint("# ID\tType\tName\tNumOfentries\tStaging\tProduction"))
	for _, singleList := range *netLists {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t%v\t%s\t%s", singleList.UniqueID, singleList.Type, singleList.Name, singleList.NumEntries, singleList.StagingActivationStatus, singleList.ProductionActivationStatus))
	}

	w.Flush()

}

// tablePrintNetworkList is responsible for pretty print of our network list
//
// cmd_list
func tablePrintNetworkList(singleList *edgegrid.AkamaiNetworkList) {

	sort.Strings(singleList.List)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)

	fmt.Fprintln(w, fmt.Sprint("# ID\tName\tNumOfentries\tStaging\tProduction"))
	fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%v\t%s\t%s", singleList.UniqueID, singleList.Name, singleList.NumEntries, singleList.StagingActivationStatus, singleList.ProductionActivationStatus))
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "# Elements")

	for iterator, ipAddress := range singleList.List {
		fmt.Fprint(w, fmt.Sprintf("%-15s\t", ipAddress))
		if (iterator+1)%8 == 0 {
			fmt.Fprintln(w, "")
		}
	}

	w.Flush()
}

// tablePrintNetworkList is responsible for pretty print of our network list
//
// cmd_list
func tablePrintNetworkListActivationStatus(singleList *edgegrid.ActivateNetworkListStatus, activationEnvironment string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, fmt.Sprint("# ID\tStatus\tActivationEnvironment\tActivationStatus\tActivationComments"))

	targetColor := color.FgYellow

	if singleList.ActivationStatus == "ACTIVE" {
		targetColor = color.FgGreen
	}

	activationColor := color.New(targetColor).SprintFunc()

	fmt.Fprintln(w, fmt.Sprintf("%s\t%v\t%s\t%s\t%s", singleList.UniqueID, singleList.Status, activationEnvironment, activationColor(singleList.ActivationStatus), singleList.ActivationComments))
	w.Flush()
}
