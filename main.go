package main

import (
	"Kaplan-Go/config"
	Constants "Kaplan-Go/constants"
	"Kaplan-Go/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	configReader := config.NewConfigReader(Constants.ConfigFile)
	configMap, err := configReader.ReadConfig()
	if err != nil {
		log.Println(err)
	}
	configuration := config.NewConfigUtil(configMap)
	appConfig := config.NewConfig(configuration)
	appRouter := router.InitializeRouter(appConfig)

	err = http.ListenAndServe(":"+appConfig.GetPort(), appRouter)
	if err != nil {
		fmt.Println("err")
	}
}
