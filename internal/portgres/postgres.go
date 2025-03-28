package portgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ConfigDB struct {
	Port     string
	Host     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(cfg ConfigDB) (*sqlx.DB, error) {
	Constr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sqlx.Open("postgres", Constr)

	if err != nil {
		log.Fatalf("cannot read db cfg:%s", err.Error())
	}

	err = db.Ping()

	if err != nil {
		logrus.Error(err.Error())
	}
	return db, nil
}
