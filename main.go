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

	ip, err := get_ext_ip(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ip)
	zoneID, err := getZoneID(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(zoneID)
}
