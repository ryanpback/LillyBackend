package main

import (
	bs "lillyAppBackend/bootstrap"
	"lillyAppBackend/helpers"
	"lillyAppBackend/httpd/handlers"
	"lillyAppBackend/models"
)

var appConfig bs.Config

func bootstrap() {
	c, err := bs.InitConfig(helpers.GetEnv(isTesting))
	if err != nil {
		panic(err)
	}

	appConfig = c
	helpers.AppConfig = c
	models.AppConfig = c
	handlers.AppConfig = c
}
