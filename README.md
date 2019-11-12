# Akamai CLI for network lists
The Akamai Network List Kit is a CLI that wraps Akamai's {OPEN} APIs to let you manage network lists and their items along with activation. You can create/list/remove and search for items and lists.

Should you miss something we *gladly accept patches* :)

CLI uses custom [Akamai API client](https://github.com/apiheat/go-edgegrid)


=================

   * [Configuration and Installation](#configuration-and-installation)
      * [Credentials](#credentials)
      * [Installation](#installation)
         * [Akamai-cli ( recommended )](#akamai-cli--recommended-)
         * [Standalone](#standalone)
   * [Usage](#usage)
      * [Account switch key (ask)](#account-switch-key-ask)
      * [Functionality](#functionality)
         * [Getting a network list](#getting-a-network-list)
         * [Searching](#searching)
         * [Removing elements](#removing-elements)
         * [Creating network list](#creating-network-list)
         * [Adding elements](#adding-elements)
         * [Synchronize list](#synchronize-list)
         * [Subscribe for notifications](#subscribe-for-notifications)
         * [Activate network list](#activate-network-list)
   * [Changes](#changes)
      * [v6.0.0](#v600)
      * [v5.0.0](#v500)
      * [v4.0.0](#v400)
   * [Development](#development)

# Configuration and Installation

## Credentials
Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Tools expect proper format of sections in edgerc file which example is shown below

>*NOTE:* Default file location is *~/.edgerc*

```
[default]
client_secret = XXXXXXXXXXXX
host = XXXXXXXXXXXX
access_token = XXXXXXXXXXXX
client_token = XXXXXXXXXXXX
```

In order to change config file which is being used you can
* change it via `--config parameter` of the tool itself
* change it via env variable `export AKAMAI_EDGERC_CONFIG=/Users/jsmitsh/.edgerc`

In order to change section which is being actively used you can
* change it via `--section parameter` of the tool itself
* change it via env variable `export AKAMAI_EDGERC_SECTION=mycustomsection`

>*NOTE:* Make sure your API client do have appropriate scopes enabled

## Installation
The tool can be used as completly standalone binary or in conjuction with akamai-cli

### Akamai-cli ( recommended )

1.  Execute the following from console
    `akamai install https://github.com/apiheat/akamai-cli-netlist`

### Standalone
As part of automated releases/builds you can download latest version from the project release page

# Usage
## Account switch key (ask)
From version 5.0.0 you can now enjoy using single credentials across all of your accounts.Please make sure the scopes are set correctly and that you have correct rights across your contracts.

```
> akamai netlist  --account-switch-key 1123-ACS  search --searchPattern "some-Pattern_" | jq '.[] | { name, type}'

> akamai netlist  --account-switch-key 1-ABC782  get by-name --name some-list-name

> akamai netlist  -ask 1-8282AD  get by-name --name other_name_of_list
```

## Functionality
In the section you will find all information regarding the usage of our tool.
### Getting a network list
Getting a network list is possible either with name of the netlist or with the id.
* get a specific network list by unique Id
   ```
   akamai netlist get by-id --id 1234_UNIQID
   ```

* get a specific network list by name
   ```
   akamai netlist get by-name --name SOME-NAME
   ```

* Get elements from a network list based on its name
   ```
   akamai netlist get by-name --name some_list_name --includeElements | jq -r '.[].list'
   ```
### Searching
You can search through network list based on name or elements which list includes
* Find all network list where it has specific element or name
   ```
   akamai netlist search --searchPattern 1.2.3.4 | jq
   ```

   ```
   akamai netlist search --searchPattern someneNa | jq
   ```

### Removing elements
* Remove element from network list
   ```
   akamai netlist items remove --id 1234_UNIQID --element 1.2.3.4
   ```

### Creating network list
* Create new network list
   ```
   akamai netlist create list --name whitelist_placeholder
   ```

### Adding elements
* Add items to network list by spectfying items to add ( comma seperated )
   ```
   akamai netlist items add --id 1234_UNIQID --items ITEM1,ITEM2,ITEM3
   ```
* Add items to network list from a file ( with IPs/CIDRs seperate per line)
   ```
   akamai netlist items add --id 1234_UNIQID --from-file /path/to/file
   ```
### Synchronize list
This new functionality ( introduced from v6 ) allows you to synchronize destination list from either local file or from another network list.
In all cases if during the sync process will recognize addresses would need to be *removed* from the target list you need to use `--force` switch to indicate that you allow for that action.
* Synchronize from file
   ```
   akamai netlist sync local --from-file /path/to/file --id-dst 1234_UNIQID
   ```
* Synchronize two network lists in Akamai
   ```
   akamai netlist sync aka --id-src 1234_DUMMYDELETE1 --id-dst 1234_UNIQID --force
   ```
### Subscribe for notifications
* Subscribe to notifications
   ```
   akamai netlist notification --networkListsIDs 1234_UNIQID --notificationRecipients rafpe@mailinator.com
   ```

* Unsubscribe from notifications
   ```
   akamai netlist notification --networkListsIDs 1234_UNIQID --notificationRecipients rafpe@mailinator.com --unsubscribe
   ```
### Activate network list
* Activate network list
   ```
   akamai netlist activate list --id 1234_UNIQID --notificationRecipients rafpe@mailinator.com

   akamai netlist activate list --id 1234_UNIQID --notificationRecipients rafpe@mailinator.com --prd
   ```
* Get activation status
   ```
   akamai netlist activate status --id 1234_UNIQID
   ```


# Changes
## v6.0.0
* introduces adding/syncing from file
* introduces sync between 2 lists
* changes to using of v6 of edgegrid client

## v5.0.0
* Use account switch key
* Improve search
* Only one list returned when searched by name ( exact match )

## v4.0.0
* Move to network lists API endpoint v2
* Use go-edgegrid client v4.X.X
* Remove support for network list API v1

# Development
In order to develop the tool with us do the following:
1. Fork repository
1. Clone it to your folder ( within *GO* path )
1. Ensure you can restore dependencies by running
   ```shell
   dep ensure -v
   ```
1. Make necessary changes
1. Make sure solution builds properly ( feel free to add tests )
   ```shell
   go build -ldflags="-s -w -X main.appVer=1.2.3 -X main.appName=$(basename `pwd`)"
   ```