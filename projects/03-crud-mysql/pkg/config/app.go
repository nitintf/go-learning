package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dbConnection, err := gorm.Open("postgres", "postgres:postgres/test-go-db")

	if err != nil {
		panic(err)
	}

	db = dbConnection
}

func GetDB() *gorm.DB {
	return db
}
