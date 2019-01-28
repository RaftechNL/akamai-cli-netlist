# Akamai CLI for network lists
The Akamai Network List Kit is a CLI that wraps Akamai's {OPEN} APIs to let you manage network lists and their items along with activation. You can create/list/remove and search for items and lists.

Should you miss something we *gladly accept patches* :)

CLI uses custom [Akamai API client](https://github.com/apiheat/go-edgegrid)

# Configuration & Installation

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

In order to change section which is being actively used you can
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

# Actions

```shell
Incorrect Usage. flag provided but not defined: -?

NAME:
    - A CLI to interact with Akamai network lists

USAGE:
    [global options] command [command options] [arguments...]

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     activate      Manages network list activation/status
     create        Creates new network list
     delete        Deletes network list ( ** REQUIRES LIST TO BE DEACTIVATED ON BOTH NETWORKS ** )
     get           List network lists objects
     items         Manages items in network lists
     notification  Manages network list subscription notifications ( SUBSCRIBE by default ) 
     search        Finds all network lists that match specific expression ( either name or network element )
     help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE   Location of the credentials FILE (default: "/Users/rafpe/.edgerc") [$AKAMAI_EDGERC_CONFIG]
   --debug value            Debug Level [$AKAMAI_EDGERC_DEBUGLEVEL]
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
flag provided but not defined: -?


```

# Using account switch key
From version 5.0.0 you can now enjoy using single credentials across all of your accounts.Please make sure the scopes are set correctly and that you have correct rights across your contracts. 

```
> akamai netlist  --account-switch-key 1123-ACS  search --searchPattern "some-Pattern_" | jq '.[] | { name, type}'

> akamai netlist  --account-switch-key 1-ABC782  get by-name --name some-list-name

> akamai netlist  --account-switch-key 1-8282AD  get by-name --name other_name_of_list
```

# Example commands
Below you may find examples on how some commands can be used together 
* get a specific network list by unique Id
   ```
   akamai netlist get by-id --id UNIQUE-ID
   ```

* get a specific network list by name
   ```
   akamai netlist get by-name --name SOME-NAME
   ```

* Get elements from a network list based on its name 
   ```
   akamai netlist get by-name --name some_list_name --includeElements | jq -r '.[].list'
   ```

* Find all network list where it has specific element or name
   ```
   akamai netlist search --searchPattern 1.2.3.4 | jq

   akamai netlist search --searchPattern someneNa | jq 
   ```

* Remove element from network list 
   ```
   akamai netlist items remove --id 1234_UNIQID --element 1.2.3.4
   ```

* Create new network list
   ```
   akamai netlist create list --name whitelist_placeholder
   ```

* Add items to network list 
   ```
   akamai netlist items add --id UNIQUE-ID --items ITEM1,ITEM2,ITEM3
   ```
* Subscribe to notifications 
   ```
   akamai netlist notification --networkListsIDs 12345_SOMENAME --notificationRecipients rafpe@mailinator.com
   ```

* Unsubscribe from notifications 
   ```
   akamai netlist notification --networkListsIDs 12345_SOMENAME --notificationRecipients rafpe@mailinator.com --unsubscribe
   ```
* Activate network list 
   ```
   akamai netlist activate list --id 12345_SOMENAME --notificationRecipients rafpe@mailinator.com

   akamai netlist activate list --id 12345_SOMENAME --notificationRecipients rafpe@mailinator.com --prd
   ```
* Get activation status
   ```
   akamai netlist activate status --id 12345_SOMENAME 
   ```


# Changes 

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