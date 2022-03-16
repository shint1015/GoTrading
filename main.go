package main

import (
	"GoTrading/bitflyer"
	"GoTrading/config"
	"GoTrading/utils"
	"fmt"
	"log"
	"time"
)

func main() {

	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	//fmt.Println(config.Config.Apikey)
	//fmt.Println(config.Config.ApiScret)

	apiClient := bitflyer.New(config.Config.Apikey, config.Config.ApiSecret)
	//fmt.Println(apiClient.GetBalance())
	ticker, _ := apiClient.GetTicker("BTC_USD")
	fmt.Println(ticker)
	fmt.Println(ticker.GetMidPrice())
	fmt.Println(ticker.DateTime())
	fmt.Println(ticker.TruncateDateTime(time.Hour))
}
