package services

import (
	"database/sql"
	"web-server/data/models"
)

var db *sql.DB

func Init(dbcon *sql.DB) {
	db = dbcon
}

func GetUsers() []models.User {
	return []models.User{}
}
