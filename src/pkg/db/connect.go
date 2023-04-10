package db

import (
	"database/sql"
	"fmt"
	"go-sql/pkg/helper"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func Connect(host string, port int) *sql.DB {
	cfg := mysql.Config{
		User:   helper.ENVS["DBUSER"],
		Passwd: helper.ENVS["DBPASS"],
		Net:    "tcp",
		Addr:   host + ":" + fmt.Sprint(port),
		DBName: helper.ENVS["DBNAME"],
	}

	log.Println("Connecting to '" + cfg.Addr + "' to table '" + cfg.DBName + "' with user '" + cfg.User + "' over '" + cfg.Net + "'")

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	log.Println("Successfully connected to MySQL database")

	return db
}
