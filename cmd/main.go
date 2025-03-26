package main

import (
	"log"
	"log/slog"
	"token-generator/config"
	"token-generator/logger"
	"token-generator/logger/sl"
	"token-generator/pkg/handlers"
	"token-generator/pkg/storage/sqlite"
	"token-generator/server"

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

	router := handlers.InitRouter(log, storage)
	srv := new(server.SRV)

	log.Info("server run on port:", slog.String("port", viper.GetString("port")))
	if err := srv.Start(viper.GetString("port"), router); err != nil {
		log.Error("cannot run server", sl.Err(err))
	}

}
