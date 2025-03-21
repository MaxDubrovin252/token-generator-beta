package main

import (
	"log"
	"token-generator/config"
	"token-generator/logger"
	"token-generator/logger/sl"
	"token-generator/pkg/storage/sqlite"

	"github.com/spf13/viper"
)

func main() {
	if err := config.SetConfig(); err != nil {
		log.Fatal(err.Error())
	}

	log, err := logger.SetLogger(viper.GetString("env"))

	if err != nil {
		log.Error("cannot init logger", sl.Err(err))
	}
	storage, err := sqlite.New(viper.GetString("storage_path"))

	if err != nil {
		log.Error("cannot connect to db", sl.Err(err))
	}

	_ = storage
	log.Info("db created")
}
