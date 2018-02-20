package main

import (
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

func config(configFile, configSection string) edgegrid.Config {
	config, err := edgegrid.Init(configFile, configSection)
	errorCheck(err)

	return config
}
