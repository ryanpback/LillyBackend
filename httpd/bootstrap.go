package main

import (
	bs "lillyAppBackend/bootstrap"
	"lillyAppBackend/helpers"
	"lillyAppBackend/httpd/handlers"
	"lillyAppBackend/models"
)

var appConfig bs.Config

func bootstrap() {
	c, err := bs.InitConfig()
	if err != nil {
		helpers.LogFatal(err)
		helpers.LogFatal(appConfig)
		panic("Someting went wrong, check your environment variables")
	}

	appConfig = c
	models.AppConfig = appConfig
	helpers.AppConfig = appConfig
	handlers.AppConfig = appConfig
}
