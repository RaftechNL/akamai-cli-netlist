package main

import (
	"log"

	common "github.com/apiheat/akamai-cli-common"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli/v2"
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
	//TODO: fix
	// common.VerifyArgumentByName(c, "id-src")
	// common.VerifyArgumentByName(c, "id-dst")

	synchronize(c.String("id-src"), c.String("id-dst"), false, c.Bool("force"))

	// resultErr := service.NetworkListErrorv2{}
	// resultErr.Title = "Sync failed"
	// resultErr.Detail = "Source list does not contain items for sync"

	// return errors.New("Source list does not have ")

	return nil
}

// syncNetListWithFile synchronizes item from src list to destination list
func syncNetListWithFile(c *cli.Context) error {
	//TODO: fix
	// common.VerifyArgumentByName(c, "from-file")
	// common.VerifyArgumentByName(c, "id-dst")

	synchronize(c.String("from-file"), c.String("id-dst"), true, c.Bool("force"))

	return nil
}

//synchronize is used to synchronize between 2 sources of IPs. If used with force option it will
//also perform removal of addresses from the target.
func synchronize(source, destination string, fromFile, force bool) {
	var ipsFromSource []string

	listNetListOptsv2 := service.ListNetworkListsOptionsv2{}
	listNetListOptsv2.IncludeElements = true

	if fromFile {
		// Get source IPs from file in local system
		sourceIPs, err := readLinesFromFile(source)
		if err != nil {
			log.Fatal(err)
		}
		ipsFromSource = sourceIPs

	} else {
		// Get source IPs from list in Akamai
		netListSrc, netlistErr := apiClient.GetNetworkList(source, listNetListOptsv2)
		if netlistErr != nil {
			log.Fatalln(netlistErr)
		}
		ipsFromSource = netListSrc.List
	}

	// Get the destination list
	netListDst, netlistErr := apiClient.GetNetworkList(destination, listNetListOptsv2)
	if netlistErr != nil {
		log.Fatalln(netlistErr)
	}

	//What is present in source list and not in destination
	diffAdd := stringsSlicesDifference(ipsFromSource, netListDst.List)

	//What is present in destination list and not in source
	diffRemove := stringsSlicesDifference(netListDst.List, ipsFromSource)

	//Safe check we will not remove all ips we have
	if len(diffRemove) > 0 && !force {
		log.Fatalf("Some IPs are present in target, but not in source. %v. Use `--force` to enable removing", diffRemove)
	}

	//Iterate and remove the values which are not present in file
	for _, IPForRemoval := range diffRemove {
		_, err := apiClient.RemoveNetworkListElement(destination, IPForRemoval)
		if err != nil {
			log.Fatal(err)
		}
	}

	//if there is nothing to add - bail out...
	if len(diffAdd) == 0 {
		return
	}

	//Add entire set from the file to the change
	syncListOpts := service.NetworkListsOptionsv2{
		List: ipsFromSource,
	}

	// Append items from src list to dst list
	netListDst, errAdd := apiClient.AddNetworkListElement(destination, syncListOpts)
	if errAdd != nil {
		log.Fatal(errAdd)
	}

	common.OutputJSON(netListDst)

	return
}
