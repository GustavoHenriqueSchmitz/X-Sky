package database

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() error {
	dsn := "host=localhost user=xsky password=xsky dbname=X-Sky port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn))
	DB = db

	if err != nil {
		return errors.New("Erro ao tentar se conectar com o banco de dados.")
	}

	return nil
}
