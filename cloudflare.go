package main

import (
	"context"
	"fmt"
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

// Pupulates dnsrecord struct and returns it.
func createRecord(Cfg *Config, ip string, name string) cloudflare.DNSRecord {
	dnsrecord := cloudflare.DNSRecord{}
	dnsrecord.Type = "A"
	dnsrecord.Name = name
	dnsrecord.Content = ip
	dnsrecord.TTL = 1 // 1 = Automatic
	return dnsrecord
}

func postDNSRecord(cfg *Config, c *CfVars, zoneID string, record cloudflare.DNSRecord) error {
	_, err := c.API.CreateDNSRecord(c.context, zoneID, record)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully added record: %v to zoneID: %s (domain: %s)\n", record, zoneID, cfg.Domain)

	return nil
}
