package config

import (
	"github.com/spf13/viper"
	"log"
)

type config struct {
	server *Server
}

var appConfig *config

func LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config, err: %v", err)
	}

	appConfig = &config{
		server: NewServerConfig(),
	}
}

func GetServerConfig() *Server {
	return appConfig.server
}
