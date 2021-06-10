package main

import (
	"context"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

type CfVars struct {
	context context.Context
	API     *cloudflare.API
}

func getCloudflareObjects(cfg *Config) *CfVars {
	c := new(CfVars)
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), cfg.Cloudflare_email)
	if err != nil {
		log.Fatal(err)
	}
	c.API = api

	c.context = context.Background()
	return c
}

// Gets zoneID which we will use to update DNS records in that zone
func getZoneID(cfg *Config, c *CfVars) (ZoneID string, err error) {

	id, err := c.API.ZoneIDByName(cfg.Domain)
	if err != nil {
		return "", err
	}
	return id, nil

}
