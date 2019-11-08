package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	common "github.com/apiheat/akamai-cli-common"
	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli"
)

func cmdAddItemsToNetlist(c *cli.Context) error {
	return addItemsToNetlist(c)
}

func addItemsToNetlist(c *cli.Context) error {

	var itemsToAdd []string

	common.VerifyArgumentByName(c, "id")

	if c.String("from-file") == "" {
		common.VerifyArgumentByName(c, "items")

		if len(c.StringSlice("items")) < 1 {
			log.Fatal("Please provide items!")

		}
	} else {
		// This is the way we access first of the params in CLI
		path := c.StringSlice("from-file")[0]

		lineItems, err := readLinesFromFile(path)
		if err != nil {
			fmt.Println(err)
		}

		for _, singleIPAddress := range lineItems {
			if singleIPAddress != "" {
				itemsToAdd = append(itemsToAdd, singleIPAddress)
			}
		}
	}

	//itemsToAdd := strings.Split(c.StringSlice("items")[0], ",")

	editListOpts := service.NetworkListsOptionsv2{
		List: itemsToAdd,
	}

	netLists, err := apiClient.AddNetworkListElement(c.String("id"), editListOpts)

	if err != nil {
		return err
	}

	common.OutputJSON(netLists)

	return nil

}

// readLinesFromFile reads a whole file into memory
// and returns a slice of its lines.
func readLinesFromFile(path string) ([]string, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Print("File Does Not Exist: ")
		}
		fmt.Println(err)
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
