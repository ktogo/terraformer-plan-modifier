package main

import (
	"log"

	"github.com/ktogo/terraformer-plan-modifier/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
