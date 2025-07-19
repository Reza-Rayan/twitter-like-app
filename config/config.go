package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name    string `mapstructure:"name"`
		Port    int    `mapstructure:"port"`
		Env     string `mapstructure:"env"`
		BaseURL string `mapstructure:"base_url"`
	} `mapstructure:"app"`

	Database struct {
		Driver       string `mapstructure:"driver"`
		Name         string `mapstructure:"name"`
		MaxIdleConns int    `mapstructure:"max_idle_conns"`
		MaxOpenConns int    `mapstructure:"max_open_conns"`
	} `mapstructure:"database"`

	Monitoring struct {
		Enabled bool   `mapstructure:"enabled"`
		Path    string `mapstructure:"path"`
	} `mapstructure:"monitoring"`
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	AppConfig = &Config{}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
}
