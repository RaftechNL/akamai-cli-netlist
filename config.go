package main

import (
	"fmt"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/fatih/color"
	"github.com/go-ini/ini"
	"github.com/urfave/cli"	
)

func config(configFile, configSection string) {
    cfg, err := ini.Load(configFile)
    if err != nil {
        color.Set(color.FgRed)
        fmt.Printf("'%s' does not exist. Please run '%s --config Your_Configuration_File'...\n", configFile, appName)
        color.Unset()
    } else {
        _, err = cfg.GetSection(configSection)
        if err != nil {
            color.Set(color.FgRed)
            fmt.Printf("Section '%s' does not exist in %s. Please run '%s --section <section-name> or define AKAMAI_EDGERC_SECTION env var' ...\n", configSection, configFile, appName)
            color.Unset()
        } else {
            var errEdge error
            // edgegrid.Init cannot handle wrong section. Should we add the check?
            edgeConfig, errEdge = edgegrid.Init(configFile, configSection)

            if errEdge != nil {
                color.Set(color.FgRed)
                fmt.Println("Error with section found, please check that all fields present")
                color.Unset()
                cli.NewExitError(errEdge.Error(), 1)
            }
        }
    }

    return
}