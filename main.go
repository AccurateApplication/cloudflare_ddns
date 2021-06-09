package main

import (
	"fmt"
)

// ipify returns external IP as json
const url = "https://api.ipify.org?format=json"

func main() {

	ip, err := get_ext_ip(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
	getZoneID()
}
