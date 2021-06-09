package main

import (
	"flag"
	"fmt"
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
	fmt.Println(config)

	ip, err := get_ext_ip(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
	getZoneID()
}
