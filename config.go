package main

import (
	"fmt"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/fatih/color"
	"github.com/go-ini/ini"
)

func config(configFile, configSection string) (edgegrid.Config, error) {
	cfg, err := ini.Load(configFile)
	if err != nil {
		color.Set(color.FgRed)
		fmt.Printf("'%s' does not exist. Please run '%s --config Your_Configuration_File'...\n", configFile, appName)
		color.Unset()
	} else {
		_, err = cfg.GetSection(configSection)
		if err != nil {
			color.Set(color.FgRed)
			fmt.Printf("Section '%s' does not exist in %s. Please run '%s --section Your_Section_Name or define AKAMAI_EDGERC_SECTION env var' ...\n", configSection, configFile, appName)
			color.Unset()
		}
	}

	config, errEdge := edgegrid.Init(configFile, configSection)

	return config, errEdge
}
