package main

import (
	"gateway/api"
	"gateway/config"
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := api.NewApp()
	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
