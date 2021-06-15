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

}
