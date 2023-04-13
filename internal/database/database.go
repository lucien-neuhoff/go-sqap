package database

import (
	"database/sql"
	"fmt"
	"go-sqap/internal/config"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase(cfg config.Config) *sql.DB {
	db_cfg := mysql.Config{
		User:   cfg.DBUsername,
		Passwd: cfg.DBPassword,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort),
		DBName: cfg.DBName,
	}

	log.Printf("Connecting to '%s/%s' on table '%s' over '%s'.", db_cfg.Addr, db_cfg.User, db_cfg.DBName, db_cfg.Net)

	var err error
	DB, err = sql.Open("mysql", db_cfg.FormatDSN())
	if err != nil {
		log.Fatal("Error while connecting to DB: ", err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal("Error while pinging DB: ", err)
	}

	return DB
}
