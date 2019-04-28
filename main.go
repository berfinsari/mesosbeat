package main

import (
	"os"

	"github.com/berfinsari/mesosbeat/cmd"

	_ "github.com/berfinsari/mesosbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
