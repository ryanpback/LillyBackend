package testhelpers

import (
	"lillyAppBackend/bootstrap"
	"lillyAppBackend/helpers"
)

// BootstrapTestConfig will return a db instance
func BootstrapTestConfig() bootstrap.Config {
	c, err := bootstrap.InitConfig(helpers.GetEnv(isTesting))
	if err != nil {
		panic(err.Error())
	}

	return c
}
