package main

import (
	"log"
	"strings"

	service "github.com/apiheat/go-edgegrid/v6/service/netlistv2"
	"github.com/urfave/cli/v2"
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
	//TODO: fix
	// common.VerifyArgumentByName(c, "networkListsIDs")
	// common.VerifyArgumentByName(c, "notificationRecipients")

	if len(c.StringSlice("notificationRecipients")) < 1 {
		log.Fatal("Please provide notificationRecipients!")

	}
	notificationRecipients := strings.Split(c.StringSlice("notificationRecipients")[0], ",")

	if len(c.StringSlice("networkListsIDs")) < 1 {
		log.Fatal("Please provide networkListsIDs!")

	}
	networkListsIDs := strings.Split(c.StringSlice("networkListsIDs")[0], ",")

	networkListSubscription := service.NetworkListSubscription{
		Recipients: notificationRecipients,
		UniqueIds:  networkListsIDs,
	}

	notificationAction := service.Subscribe
	if c.Bool("unsubscribe") {
		notificationAction = service.Unsubscribe
	}

	netlistErr := apiClient.NetworkListNotification(notificationAction, networkListSubscription)
	if netlistErr != nil {
		return netlistErr
	}
	return nil
}
