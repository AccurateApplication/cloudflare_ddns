package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Ip struct {
	Ip string `json:"ip"`
}

func get_ext_ip(url string) (ip string, err error) {
	r, err := http.Get(url)
	if err != nil {
		log.Printf("Trouble getting %s, error: %v", url, err.Error())
		return "error", err
	}
	decoder := json.NewDecoder(r.Body)
	var i Ip
	err = decoder.Decode(&i)
	return i.Ip, nil
}
