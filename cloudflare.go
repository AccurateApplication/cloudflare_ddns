package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func getZoneID() {
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		log.Fatal(err)
	}

	id, err := api.ZoneIDByName("domain")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

}
