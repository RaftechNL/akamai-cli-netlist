package main

import (
	"log"
	"strings"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	"github.com/urfave/cli"
)

// cmdNotificationManagement is used by cli to execute notification management
//
// cmd_search
func cmdNotificationManagement(c *cli.Context) error {
	return notificationManagement(c)
}

// notificationManagement execute client API call to notification subscription modification
//
// cmd_search
func notificationManagement(c *cli.Context) error {
	common.VerifyArgumentByName(c, "networkListsIDs")
	common.VerifyArgumentByName(c, "notificationRecipients")

	if len(c.StringSlice("notificationRecipients")) < 1 {
		log.Fatal("Please provide notificationRecipients!")

	}
	notificationRecipients := strings.Split(c.StringSlice("notificationRecipients")[0], ",")

	if len(c.StringSlice("networkListsIDs")) < 1 {
		log.Fatal("Please provide networkListsIDs!")

	}
	networkListsIDs := strings.Split(c.StringSlice("networkListsIDs")[0], ",")

	networkListSubscription := edgegrid.NetworkListSubscription{
		Recipients: notificationRecipients,
		UniqueIds:  networkListsIDs,
	}

	notificationAction := edgegrid.Subscribe
	if c.Bool("unsubscribe") {
		notificationAction = edgegrid.Unsubscribe
	}

	_, netlistErr := apiClient.NetworkListsv2.NetworkListNotification(notificationAction, networkListSubscription)
	if netlistErr != nil {
		return netlistErr
	}
	return nil
}
