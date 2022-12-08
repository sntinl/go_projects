package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DSN = "host=localhost user=postgres password=abcABC123 dbname=store port=5432 sslmode=disable"

var DB *gorm.DB

func DbConnnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB CONNECTED")
	}

}
