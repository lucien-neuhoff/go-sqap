package main

import (
	api_server "go-sql/pkg/api"
	db "go-sql/pkg/db"
	helper "go-sql/pkg/helper"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	getEnvs()

	db_port, err := strconv.Atoi(helper.ENVS["DBPORT"])
	if err != nil {
		log.Fatal(err)
		db_port = 3306
	}
	helper.DB = db.Connect(helper.ENVS["DBHOST"], db_port)

	api_port, err := strconv.Atoi(helper.ENVS["APIPORT"])
	if err != nil {
		log.Fatal(err)
		api_port = 8080
	}
	api_server.Start(helper.ENVS["APIHOST"], api_port)
}

func getEnvs() {
	var err error
	helper.ENVS, err = godotenv.Read(".env")

	if err != nil {
		log.Fatal(err)
	}
}
