package main

import (
	"encoding/json"
	"net/http"
)

type Ip struct {
	Ip string `json:"ip"`
}

func get_ext_ip(url string) (ip string, err error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}
	decoder := json.NewDecoder(r.Body)
	var i Ip
	err = decoder.Decode(&i)
	return i.Ip, nil
}
