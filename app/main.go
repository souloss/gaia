package main

import (
	"log"

	api "github.com/witchc/gaia/api"
)

func main() {
	if err := api.StartService(); err != nil {
		log.Fatal(err)
	}
}
