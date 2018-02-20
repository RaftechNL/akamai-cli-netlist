# Akamai CLI for network lists
*NOTE:* This tool is intended to be installed via the Akamai CLI package manager, which can be retrieved from the releases page of the [Akamai CLI](https://github.com/akamai/cli) tool.

The Akamai Network List Kit is a set of go libraries that wraps Akamai's {OPEN} APIs to let you manage network lists and their items along with activation. You can create/list/remove and search for items and lists.

This tool have been created with idea of using *akamai cli*. It supports most of the methods provided by API of Akamai. Should you miss something we gladly *accept patches* :)

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
    `akamai install https://github.com/RafPe/akamai-cli-networklist`
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



## Commands

### App overview
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
   --config FILE, -c FILE   Location of the credentials FILE (default: "/Users/username/.edgerc") [$AKAMAI_EDGERC_CONFIGFILE]
   --no-color               Disable color output
   --output value           Prints json output (raw) or table (default: "table")
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```