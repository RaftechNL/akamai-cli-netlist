# Akamai CLI for network lists
*NOTE:* This tool is intended to be installed via the Akamai CLI package manager, which can be retrieved from the releases page of the [Akamai CLI](https://github.com/akamai/cli) tool.

The Akamai Network List Kit is a set of go libraries that wraps Akamai's {OPEN} APIs to let you manage network lists and their items along with activation. You can create/list/remove and search for items and lists.

This tool have been created with idea of using *akamai cli*. It supports most of the methods provided by API of Akamai. Should you miss something we gladly *accept patches* :)

It uses custom [Akamai API client](https://github.com/RafPe/go-edgegrid)

<!--ts-->
   * [Akamai CLI for network lists](#akamai-cli-for-network-lists)
   * [Configuration &amp; Installation](#configuration--installation)
      * [API Credentials and sections](#api-credentials-and-sections)
      * [Installation](#installation)
         * [Via akamai-cli ( recommended )](#via-akamai-cli--recommended-)
         * [Standalone](#standalone)
      * [App overview](#app-overview)
         * [General](#general)
         * [Get](#get)
         * [Create](#create)
         * [Activate](#activate)
         * [Search](#search)
         * [Remove](#remove)
      * [Credits](#credits)

<!--te-->
# Configuration & Installation

## API Credentials and sections
Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Tools expect proper format of sections in .edgerc which looks as follow

```
[default]
client_secret = XXXXXXXXXXXX
host = XXXXXXXXXXXX
access_token = XXXXXXXXXXXX
client_token = XXXXXXXXXXXX
```

In order to change section which is being actively used you can
* change via `--section parameter` of the tool itself
* change via env variable `export AKAMAI_EDGERC_SECTION=mycustomsection`

Make sure your API client do have approiate scopes enabled to manage network lists

## Installation
Available in two different ways.With akamai-cli toolkit or as a standalone version

### Via akamai-cli ( recommended )

1.  Execute the following from console     
    `akamai install https://github.com/RafPe/akamai-cli-netlist`

### Standalone
To compile it from source, you will need Go 1.9 or later, and the [Glide](https://glide.sh) package manager installed:
1. Fetch the package:
   `go get https://github.com/RafPe/akamai-cli-networklist`
1. Change to the package directory:
   `cd $GOPATH/src/github.com/RafPe/akamai-cli-networklist`
1. Install dependencies using Glide:
   `glide install`
1. Compile the binary:
   `go build -ldflags="-s -w -X main.version=X.X.X" -o akamai-netlist`



## App overview

### General
```shell
NAME:
   akamai-netlist - A CLI to interact with Akamai Network lists

USAGE:
   akamai-netlist [global options] command [command options] [arguments...]

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     activate  Manage network list activation
     create    Creates network list
     get       List network lists objects
     remove    removes network list/items
     search    search by expression
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE   Location of the credentials FILE (default: "/Users/username/.edgerc") [$AKAMAI_EDGERC_CONFIG]
   --no-color               Disable color output
   --output value           Prints json output (raw) or table (default: "table")
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```
### Get
This main command allows you to exexute `get` actions on Akamai.

* Getting all network lists

    ```shell
    > $ akamai netlist get all
    # ID                            Name                                    NumOfentries
    36547_RPIENIAZEKTEST1           rpieniazek_test_1                       2
    36543_APITEST                   api_test                                1
    36542_APITEST                   api_test                                0
    36541_APITEST                   api_test                                0
    36540_APITEST                   api_test                                0
    36539_APITEST                   api_test                                0
    36538_APITEST                   api_test                                0
    36537_APITEST                   api_test                                0
    36536_APITEST                   api_test                                0
    36535_APITEST                   api_test                                0
    ```
* Getting single list

    ```shell
    > $ akamai netlist get list --id 36547_RPIENIAZEKTEST1
    # ID                    Name                NumOfentries
    36547_RPIENIAZEKTEST1   rpieniazek_test_1   2
    ```

    ```shell
    > $ akamai netlist get list --id 36547_RPIENIAZEKTEST1 --extended
    # ID                    Name                NumOfentries   Staging   Production
    36547_RPIENIAZEKTEST1   rpieniazek_test_1   2              ACTIVE    MODIFIED
    ```
* Getting items from the list

    ```shell
    > $ akamai netlist get list --includeElements --id 38492_DUMMYDELETE1
    # ID                 Name             NumOfentries   Staging   Production
    38492_DUMMYDELETE1   dummy_delete_1   2

    # Elements
    1.2.3.4/32        5.6.7.8/32
    ```
* Passing output to `jq` is also easy and possible 

    ```shell
    > $ akamai netlist --output json get list --includeElements --id 38492_DUMMYDELETE1 | jq
    {
        "name": "dummy_delete_1",
        "type": "IP",
        "list": [
            "1.2.3.4/32",
            "5.6.7.8/32"
        ],
        "links": [
            {
                "rel": "get 38492_DUMMYDELETE1",
                "href": "/network-list/v1/network_lists/38492_DUMMYDELETE1"
            }
    ],
        "unique-id": "38492_DUMMYDELETE1",
        "sync-point": 0,
        "numEntries": 2
    }
    ```


### Create
Used for creating new network lists of network list items

* Create `IP` network list
    ```shell
    > $ akamai netlist create list --name="test-readme"
    ok
    ```

* Create `GEO` network list
    ```shell
    > $ akamai netlist create list --name="test-readme" --type GEO
    ok
    ```

* Create new item in list ( items needs to be comma separated )
    ```shell
    > $ akamai netlist create item --id 36547_RPIENIAZEKTEST1 --items 1.2.3.4/32,2.3.4.5/24
    Elements successfully appended to the list
    ```

### Activate
Used for view of activation status or activating a list itself

```shell
> $ akamai netlist activate list --help
NAME:
   akamai-netlist activate list - activates network list

USAGE:
   akamai-netlist activate list [command options] [arguments...]

OPTIONS:
   --id value                      list unique-id
   --ticket-id value               ticket for this activation (default: "na")
   --comments value                ticket for this activation (default: "created via akamai-cli-networklist")
   --NotificationRecipients value  actNotificationRecipients to be included in activation email
   --prd                           activate on production

```

* Activate a network list

    ```shell
    > $ akamai netlist activate list --id 36547_RPIENIAZEKTEST1 --ticket-id 1234 --comments 'readme'
    ok
    ```

* Activate a network list in `production`

    ```shell
    > $ akamai netlist activate list --id 36547_RPIENIAZEKTEST1 --ticket-id 1234 --comments 'readme' --prd
    ok
    ```

* View activation status 

    ```shell
    > $ akamai netlist activate status  --id 36547_RPIENIAZEKTEST1
    # ID                    Status   ActivationEnvironment   ActivationStatus              ActivationComments
    36547_RPIENIAZEKTEST1   200      staging                 PENDING_ACTIVATION   readme
    ```

* View activation status in `production`

    ```shell
    > $ akamai netlist activate status  --id 36547_RPIENIAZEKTEST1 --prd
    # ID                    Status   ActivationEnvironment   ActivationStatus              ActivationComments
    36547_RPIENIAZEKTEST1   200      production              PENDING_ACTIVATION   readme-for-prd
    ```

### Search 
Used to find where specific items are included in network lists

```
> $ akamai netlist search --help
NAME:
   akamai-netlist search - search by expression

USAGE:
   akamai-netlist search [command options] [arguments...]

OPTIONS:
   --item value         item to search for
   --extended           returns more verbose data such as creation date and activation status
   --includeDeprecated  includes network lists that have been deleted
   --listType value     filters by the network list type [ IP | GEO ] (default: "IP")
```

* Search network lists containing element

    ```shell
    > $ akamai netlist search --item 1.2.3.4/32
    # ID                    Name                NumOfentries
    36547_RPIENIAZEKTEST1   rpieniazek_test_1   4
    ```

### Remove 
Allows to remove item in list or list itself.

* Remove item from list

    ```shell
    > $ akamai netlist remove item --id 36547_RPIENIAZEKTEST1 --item 1.2.3.4/32
    Element 1.2.3.4/32 successfully deleted from list
    ```

* Remove list ( this requires list to be inactive/deactivated )
    ```shell
    > $ akamai netlist remove list --id 36547_RPIENIAZEKTEST1
    The network list is not inactive in the staging network and cannot be deprecated
    ```

## Credits
* [Petr](https://github.com/partamonov) - for being the mentor on Golang :)
