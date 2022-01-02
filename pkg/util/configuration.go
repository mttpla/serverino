//Copyright 2021 Matteo Paoli - mttpla@gmail.com

package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

var (
	Configuration Config
)

func SetDefault() {
	viper.SetDefault("PORT", "8080")
}

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./cmd/")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	SetDefault()

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		Error.Println("Failed to read the config file! Keep Default values.")
	}
	viper.Unmarshal(&Configuration)
}
