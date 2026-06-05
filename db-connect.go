package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connect() *gorm.DB {
	var db *gorm.DB
	var dns string
	var err error

	dns = "host=127.0.0.1 user=root password=root123 dbname=go_inventory sslmode=disable"
	db, err = gorm.Open(postgres.Open(dns))

	if err != nil {
		panic(err)
	}
	return db
}
