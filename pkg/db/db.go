package db

import (
	"msgo-account/pkg/db/models"
  "github.com/jmoiron/sqlx"
)

type DB struct {
  db *sqlx.DB 
}

type TransactionDB interface {
	CreateTransation(p *models.Transaction) error
}
