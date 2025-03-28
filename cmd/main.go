package main

import (
	"log/slog"
	"os"
	"token-generator/config"
	"token-generator/internal/handler"
	"token-generator/internal/portgres"
	"token-generator/internal/server"
	"token-generator/pkg/reposiroty"
	"token-generator/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	if err := config.InitConfig(); err != nil {
		logrus.Errorf("cannot read config:%s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Errorf("cannot set env:%s", err.Error())
	}
	db, err := portgres.NewDB(portgres.ConfigDB{
		Port:     viper.GetString("db.port"),
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})

	if err != nil {
		logrus.Errorf("cannot connect to db error:%s", err.Error())
	}
	repos := reposiroty.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)
	srv := new(server.SRV)

	logrus.Info("START", slog.String("port", viper.GetString("port")))

	if err := srv.Start(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Errorf("cannot start server error:%s", err.Error())
	}

	//storage
	//server
}
