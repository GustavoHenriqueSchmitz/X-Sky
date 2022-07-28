package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() error {
	dsn := "host=localhost user=xsky password=xsky dbname=X-Sky port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, _ := gorm.Open(postgres.Open(dsn))
	DB = db

	return nil
}
