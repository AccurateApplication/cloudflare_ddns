package main

import (
	"flag"
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

	// Get zoneID
	zoneID, err := getZoneID(config, CfVars)
	if err != nil {
		log.Fatal(err)
	}

	// List DNS records that matches subdomain in config file.
	dnsRecords, err := listDNSRecords(config, CfVars, zoneID, config.Subdomain)
	if err != nil {
		log.Println(err)
	}

	// Get ext IP
	externalIp, err := get_ext_ip(url)
	if err != nil {
		log.Fatal(err)
	}

	// Populate struct that will be added as record to CF
	subDomainRecord := createRecord(config, externalIp, config.Subdomain)
	err = checkRecords(config, CfVars, zoneID, dnsRecords, subDomainRecord, externalIp)
	if err != nil {
		log.Fatal(err)
	}

}
