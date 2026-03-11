package database

import (
	"time"

	"github.com/gofiber/fiber/v3/log"
	go_ora "github.com/sijms/go-ora/v2"
	"github.com/jmoiron/sqlx"
	"golang.project.structure/config"
)

var (
	OraDB *sqlx.DB
)

func Initialize() error {

	conf := config.LoadConfig()

	conStr := go_ora.BuildUrl(
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
		conf.DBUser,
		conf.DBPassword,
		nil)

	db, err := sqlx.Open("oracle", conStr)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(15 * time.Minute)

	OraDB = db

	log.Info("Oracle DB connection established")

	return nil

}
