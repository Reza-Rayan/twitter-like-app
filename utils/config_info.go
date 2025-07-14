package utils

import "github.com/Reza-Rayan/twitter-like-app/config"

func GetBaseURL() string {
	return config.AppConfig.App.BaseURL
}

func GetPort() int {
	return config.AppConfig.App.Port
}
