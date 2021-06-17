package main

import (
	"context"
	"fmt"
	"log"
	"os"

	cloudflare "github.com/cloudflare/cloudflare-go"
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

// Populates cf.dnsrecord struct and returns it.
func createRecord(Cfg *Config, ip string, recordName string) cloudflare.DNSRecord {
	dnsrecord := cloudflare.DNSRecord{}
	dnsrecord.Type = "A"
	dnsrecord.Name = recordName
	dnsrecord.Content = ip
	dnsrecord.TTL = 1 // 1 = Automatic
	return dnsrecord
}

func postDNSRecord(cfg *Config, c *CfVars, zoneID string, record cloudflare.DNSRecord) error {
	_, err := c.API.CreateDNSRecord(c.context, zoneID, record)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully added record. content: %s, type: %s, name: %s to zoneID: %s (domain: %s)\n", record.Content, record.Type, record.Name, zoneID, cfg.Domain)

	return nil
}

// Returns DNS records that matches recordName (example www.subdomain.domain.org)
func listDNSRecords(cfg *Config, c *CfVars, zoneID string, recordName string) ([]cloudflare.DNSRecord, error) {
	subDomainRecord := cloudflare.DNSRecord{Name: recordName}
	rec, err := c.API.DNSRecords(c.context, zoneID, subDomainRecord)
	if err != nil {
		return nil, err

	}

	return rec, nil
}

func checkRecords(cfg *Config, c *CfVars, zoneID string, cfRecords []cloudflare.DNSRecord, subDomainRecord cloudflare.DNSRecord, externalIp string) error {
	lenRecords := len(cfRecords)
	switch {
	case lenRecords == 0:
		fmt.Println("No records added for domain, will add")
		err := postDNSRecord(cfg, c, zoneID, subDomainRecord)
		if err != nil {
			log.Println(err)
			return err
		}
	case lenRecords >= 1:
		log.Printf("Found %d records matching domain. Will check", lenRecords)
		for i, r := range cfRecords {
			fmt.Printf("Checking record: %d\tcontent: %s name: %s, ID: %s\n", i, r.Content, r.Name, r.ID)
			if r.Content == externalIp {
				fmt.Printf("Keeping %s, %s\n", r.Name, r.Content)
			} else if r.Content != externalIp {

				fmt.Printf("Will delete %s, %s\n", r.Name, r.Content)
				// Delete record without matching ext IP
				err := c.API.DeleteDNSRecord(c.context, zoneID, r.ID)
				if err != nil {
					log.Println(err)
					return err
				}

				// Post DNS record with correct ext IP
				err = postDNSRecord(cfg, c, zoneID, subDomainRecord)
				if err != nil {
					log.Println(err)
					return err
				}
			}
		}
	}
	return nil
}
