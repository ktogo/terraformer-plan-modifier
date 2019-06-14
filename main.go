package main

import (
	"log"

	"github.com/ktogo/terraformer-plan-splitter/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
