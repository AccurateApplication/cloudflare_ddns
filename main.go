package main

import (
	"flag"
	"log"
	"time"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "configfile", "./config.toml", "file that will be parsed for configuration")
}

func main() {
	flag.Parse()

	config := readConfig()
	CfVars := getCloudflareObjects(config)

	log.Printf("Will check DNS record every %d minutes.\n", config.RefreshRate)

	// Get zoneID
	zoneID, err := getZoneID(config, CfVars)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Will change first time run
		curIp := "1.1.1.1"

		// Get ext IP
		tmpIp, err := get_ext_ip(config.ExtIpUrl)
		if err != nil {
			log.Panic(err)
		}

		if tmpIp != curIp {
			curIp = tmpIp
			// Populate struct that will be added as record to CF
			subDomainRecord := createRecord(config, curIp, config.Subdomain)

			// List DNS records that matches subdomain in config file.
			dnsRecords, err := listDNSRecords(config, CfVars, zoneID, config.Subdomain)
			if err != nil {
				log.Println(err)
			}

			err = checkRecords(config, CfVars, zoneID, dnsRecords, subDomainRecord, curIp)
			if err != nil {
				log.Fatal(err)
			}
		}

		time.Sleep(time.Duration(config.RefreshRate) * time.Minute)

	}

}
