package configs

import "os"

type GlobalConfig struct {
	AppPort string
	AppHost string
}

var GlobalConfiguration GlobalConfig

func SetGlobalConfiguration() {
	GlobalConfiguration.AppPort = os.Getenv("APP_PORT")
	GlobalConfiguration.AppHost = os.Getenv("APP_HOST")
}
