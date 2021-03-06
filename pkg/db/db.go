package db

import (
	"log"
	"msgo-account/pkg/repository"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type GeneralDB interface {
	Open() error
	Close() error
	TransactionDB
	AccountDB
  UserDB
  CategoryDB
  CurrencyDB
  RateDB
  BudgetDB
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Cannot open dotenv file: %s", err.Error())
	}

	pg, err := repository.InitDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println("Connected to DB")
	d.db = pg
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
