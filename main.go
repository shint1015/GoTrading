package main

import (
	"GoTrading/config"
	"GoTrading/utils"
	"log"
)

func main() {

	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	//fmt.Println(config.Config.Apikey)
	//fmt.Println(config.Config.ApiScret)
}
