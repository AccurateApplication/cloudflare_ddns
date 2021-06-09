package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func getZoneID(cfg Config) {
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), cfg.Cloudflare_email)
	if err != nil {
		log.Fatal(err)
	}

	id, err := api.ZoneIDByName(cfg.Domain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

}
