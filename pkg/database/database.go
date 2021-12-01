package database

import (
	"fmt"
	"go-echo-restful/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DBConn *gorm.DB

func UseDatabase() {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s", config.Cfg.DatabaseHost, config.Cfg.DatabaseUser, config.Cfg.DatabasePassword, config.Cfg.DatabaseName)

	var err error

	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database.", err)
	}
}
