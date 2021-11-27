package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	AppHost    string `mapstructure:"app_host"`
	AppPort    int    `mapstructure:"app_port"`
	PgHost     string `mapstructure:"pg_host"`
	PgPort     int    `mapstructure:"pg_port"`
	PgUsername string `mapstructure:"pg_username"`
	PgPassword string `mapstructure:"pg_password"`
	PgDbname   string `mapstructure:"pg_dbname"`
}

func GetConfiguration() *AppConfig {
	var config = AppConfig{
		PgHost:     "127.0.0.1",
		PgPort:     5432,
		PgUsername: "youruser",
		PgPassword: "yourpassword",
		PgDbname:   "yourdb",
	}

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error reading file config, set config to default")
		return &config
	}

	defaConfig := config
	if err := viper.Unmarshal(&config); err != nil {
		log.Println("Error unmarshal file config, set config to default")
		return &defaConfig
	}

	return &config
}
