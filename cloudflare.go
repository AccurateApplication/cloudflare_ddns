package main

import (
	"os"

	"github.com/cloudflare/cloudflare-go"
)

// Gets zoneID which we will use to update DNS records in that zone
func getZoneID(cfg Config) (ZoneID string, err error) {

	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), cfg.Cloudflare_email)
	if err != nil {
		return "", err
	}

	id, err := api.ZoneIDByName(cfg.Domain)
	if err != nil {
		return "", err
	}
	return id, nil

}
