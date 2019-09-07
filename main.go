package main

import (
	"os"

	"github.com/berfinsari/mesosbeat/cmd"

	// Make sure all your modules and metricsets are linked in this file
	_ "github.com/berfinsari/mesosbeat/include"
)


func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
