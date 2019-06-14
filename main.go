package main

import (
	"log"
	"os"

	"github.com/ktogo/terraformer-plan-splitter/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
