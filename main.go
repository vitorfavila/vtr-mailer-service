package main

import (
	"example/vtr-mailer-service/application"
	"example/vtr-mailer-service/config"
	"example/vtr-mailer-service/router"
	"log"
)

func main() {
	err := config.SetupDotEnv()
	if err != nil {
		log.Fatal("Failed to Setup ENV")
	}

	app, err := application.SetupApplication()
	if err != nil {
		log.Fatal(err.Error())
	}

	routerConfig := router.SetupRouter(app)

	routerConfig.Run()
}
