package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		Dsn string
	}

	Jwt struct {
		Secret string
	}
}

func LoadConfig(configName string, configType string) *Config {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(".")

	viper.SetDefault("server.port", "8080")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	cfg := new(Config)

	if err := viper.Unmarshal(cfg); err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	return cfg
}
