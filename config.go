package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Domain           string `toml:"domain"`
	Subdomain        string `toml:"subdomain"`
	Cloudflare_email string `toml:"cloudflare_email"`
}

func readConfig() *Config {
	// var c Config
	c := new(Config)
	_, err := toml.DecodeFile(configFile, &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
