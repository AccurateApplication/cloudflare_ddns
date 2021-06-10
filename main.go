package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
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

	// Create API client which we will used in cloudflare functions
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), config.Cloudflare_email)
	if err != nil {
		log.Fatal(err)
	}

	ip, err := get_ext_ip(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ip)

	zoneID, err := getZoneID(config, api)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(zoneID)

}
