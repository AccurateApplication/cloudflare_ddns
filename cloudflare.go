package main

import (
	"github.com/cloudflare/cloudflare-go"
)

// Gets zoneID which we will use to update DNS records in that zone
func getZoneID(cfg Config, api *cloudflare.API) (ZoneID string, err error) {

	id, err := api.ZoneIDByName(cfg.Domain)
	if err != nil {
		return "", err
	}
	return id, nil

}
