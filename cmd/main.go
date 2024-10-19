package main

import (
	"database/sql"
	"log"
	"test-project/cmd/api"
	"test-project/config"
	"test-project/db"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewSQLStorage(mysql.Config{
		User:                 config.ENVS.DBUser,
		Passwd:               config.ENVS.DBPassword,
		Addr:                 config.ENVS.DBAddress,
		DBName:               config.ENVS.DBName,
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)
	server := api.NewAPIServer(":3000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Successfully connected!")
}
