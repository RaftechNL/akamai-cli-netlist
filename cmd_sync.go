package main

import (
	"errors"
	"fmt"
	"log"

	common "github.com/apiheat/akamai-cli-common"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli"
)

// cmdSyncNetListID is used by cli to sync items between source and target list
func cmdSyncNetListID(c *cli.Context) error {
	return syncNetListbyID(c)
}

// cmdsyncNetListWithFile is used by cli to sync items between local file and target akamai network list
func cmdsyncNetListWithFile(c *cli.Context) error {
	return syncNetListWithFile(c)
}

// syncNetListbyID synchronizes item from src list to destination list
func syncNetListbyID(c *cli.Context) error {
	common.VerifyArgumentByName(c, "id-src")
	common.VerifyArgumentByName(c, "id-dst")

	listNetListOptsv2 := service.ListNetworkListsOptionsv2{}
	listNetListOptsv2.IncludeElements = true

	netListSrc, netlistErr := apiClient.GetNetworkList(c.String("id-src"), listNetListOptsv2)
	if netlistErr != nil {
		return netlistErr
	}

	// if we have items to sync in our source list
	// we proceed
	if len(netListSrc.List) > 0 {
		// Assign the items from src list to options obj
		syncListOpts := service.NetworkListsOptionsv2{
			List: netListSrc.List,
		}

		// Append items from src list to dst list
		netListDst, err := apiClient.AddNetworkListElement(c.String("id-dst"), syncListOpts)
		if err != nil {
			return err
		}

		common.OutputJSON(netListDst)

	}

	resultErr := service.NetworkListErrorv2{}
	resultErr.Title = "Sync failed"
	resultErr.Detail = "Source list does not contain items for sync"

	return errors.New("Source list does not have ")
}

// syncNetListWithFile synchronizes item from src list to destination list
func syncNetListWithFile(c *cli.Context) error {
	common.VerifyArgumentByName(c, "from-file")
	common.VerifyArgumentByName(c, "id-dst")

	sourceIPs, err := readLinesFromFile(c.String("from-file"))
	if err != nil {
		log.Fatal(err)
	}

	//get the items from destination list
	listNetListOptsv2 := service.ListNetworkListsOptionsv2{}
	listNetListOptsv2.IncludeElements = true

	netListDst, netlistErr := apiClient.GetNetworkList(c.String("id-dst"), listNetListOptsv2)
	if netlistErr != nil {
		return netlistErr
	}

	//What is present in Local file and not in Akamai
	diffAdd := stringsSlicesDifference(sourceIPs, netListDst.List)
	log.Printf("In local file but not in akamai .... %v", diffAdd)

	//What is present in Akamai and not in local list - we get the diff
	diffRemove := stringsSlicesDifference(netListDst.List, sourceIPs)

	//Safe check we will not remove all ips we have
	if len(diffRemove) > 0 && !c.Bool("force") {
		log.Fatalf("Some IPs are present in Akamai, but not in file. %v. Use `--force` to enable removing", diffRemove)
	}

	//Iterate and remove the values which are not present in file
	for _, IPForRemoval := range diffRemove {
		log.Println("Removing.... " + IPForRemoval)
		_, err := apiClient.RemoveNetworkListElement(c.String("id-dst"), IPForRemoval)
		if err != nil {
			log.Fatal(err)
		}
	}

	//if there is nothing to add - bail out...
	if len(diffAdd) == 0 {
		return nil
	}

	//Add entire set from the file to the change
	syncListOpts := service.NetworkListsOptionsv2{
		List: sourceIPs,
	}

	log.Println(fmt.Sprintf("Adding to the list... %v", sourceIPs))

	// Append items from src list to dst list
	netListDst, errAdd := apiClient.AddNetworkListElement(c.String("id-dst"), syncListOpts)
	if errAdd != nil {
		log.Fatal(errAdd)
	}

	common.OutputJSON(netListDst)

	return nil
}
