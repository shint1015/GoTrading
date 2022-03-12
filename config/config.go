package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Apikey   string
	ApiScret string
	LogFile  string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Apikey:   cfg.Section("bitflyer").Key("api_key").String(),
		ApiScret: cfg.Section("bitflyer").Key("api_secret").String(),
		LogFile:  cfg.Section("gotrading").Key("log_file").String(),
	}
}
