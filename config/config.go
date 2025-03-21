package config

import "github.com/spf13/viper"

func SetConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("local")
	return viper.ReadInConfig()
}
