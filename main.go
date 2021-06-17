package main

import (
	"flag"
	"fmt"
	"log"
)

// ipify returns external IP as json
const url = "https://api.ipify.org?format=json"

var configFile string

func init() {
	flag.StringVar(&configFile, "configfile", "./config.toml", "file that will be parsed for configuration")
}

func main() {
	flag.Parse()

	config := readConfig()
	CfVars := getCloudflareObjects(config)

	externalIp, err := get_ext_ip(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(externalIp)

	zoneID, err := getZoneID(config, CfVars)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(zoneID)

	subDomainRecord := createRecord(config, externalIp, config.Subdomain)
	err = postDNSRecord(config, CfVars, zoneID, subDomainRecord)
	if err != nil {
		log.Fatal(err)
	}

	// Populate struct that will be added as record to CF
	subDomainRecord := createRecord(config, externalIp, config.Subdomain)

	// TODO: Add this as a function
	lenRecords := len(dnsRecords)
	switch {
	case lenRecords == 0:
		fmt.Println("No records added for domain, will add")
		err = postDNSRecord(config, CfVars, zoneID, subDomainRecord)
		if err != nil {
			log.Println(err)
		}
	case lenRecords >= 1:
		log.Printf("Found %d records matching domain. Will check", lenRecords)
		for i, r := range dnsRecords {
			fmt.Printf("Record: %d\tcontent: %s name: %s, ID: %s\n", i, r.Content, r.Name, r.ID)
			if r.Content != externalIp {

				// Clear all records
				err = CfVars.API.DeleteDNSRecord(CfVars.context, zoneID, r.ID)
				if err != nil {
					log.Println(err)
				}

				// Post DNS record with correct ext IP
				err = postDNSRecord(config, CfVars, zoneID, subDomainRecord)
				if err != nil {
					log.Println(err)
				}
			} else if r.Content == externalIp {
				fmt.Println("Found record matching external IP. Will keep.")
			}
		}
	}

}
