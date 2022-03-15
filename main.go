package main

import (
	"GoTrading/bitflyer"
	"GoTrading/config"
	"GoTrading/utils"
	"fmt"
	"log"
)

func main() {

	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	//fmt.Println(config.Config.Apikey)
	//fmt.Println(config.Config.ApiScret)

	apiClient := bitflyer.New(config.Config.Apikey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())

}
