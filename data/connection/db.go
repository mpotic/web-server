package dbConnection

import (
	"database/sql"
	"log"
	"web-server/config"
)

var db *sql.DB

func Connect() *sql.DB {
	var config *config.Config = config.GetConfig()
	var err error
	db, err = sql.Open("postgres", config.Database.ConnectionString)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}

func Close() {
	db.Close()
}

func Init() {
	createTaskTable()
}

func createTaskTable() error {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS tasks(
			id SERIAL PRIMAL KEY,
			username VARCHAR(50) NOT NULL,
			email VARCAHR(100) NOT NULL
		);
	`)

	if err != nil {
		log.Fatal("Error trying to create a task table! ", err)
		return err
	}

	return nil
}
