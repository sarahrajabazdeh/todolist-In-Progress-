package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/sarahrajabazdeh/todolist/config"
)

type DatabaseInterface interface {
}

type Database struct {
	DB *sql.DB
}

func CreateDatabase() Database {
	var err error

	db := Database{}

	// create the connection string
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Config.Database.DbAddr, config.Config.Database.DbPort, config.Config.Database.DbUser, config.Config.Database.DbPass, config.Config.Database.DbName)

	db.DB, err = sql.Open(config.Config.Database.DbType, connectionString)
	if err != nil {
		log.Println("error db connection")
		log.Fatal(err)
	}
	// check connection
	err = db.DB.Ping()
	if err != nil {
		log.Println("Cannot connect to the db")
		log.Fatal(err)
	}
	log.Println("Connected to database ", config.Config.Database.DbName)

	return db
}
