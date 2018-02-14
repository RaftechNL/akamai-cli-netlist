package main

import (
	"log"

	"github.com/fatih/color"
)

func errorCheck(e error) {
	if e != nil {
		color.Set(color.FgRed)
		log.Fatal(e)
		color.Unset()
	}
}
