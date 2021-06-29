package db

import (
	"msgo-account/pkg/db/models"
)

type TransactionDB interface {
	CreateTransation(p *models.Transaction) error
}
