package db

import (
	"fmt"
	"msgo-account/pkg/db/models"
)

type TransactionDB interface {
	GetTransactions() ([]*models.Transaction, error)
	CreateTransaction(t *models.Transaction) error
	DeleteTransaction(t *models.DeleteTransaction) error
	UpdateTransaction(t *models.Transaction) error
}

func (d *DB) GetTransactions() ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := d.db.Select(&transactions, getTransactionsSchema)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (d *DB) CreateTransaction(t *models.Transaction) error {
	fmt.Println(t.UserId, t.CategoryId, t.Amount, t.AccountId, t.Description)
	res, err := d.db.Exec(insertTransactionSchema, t.UserId, t.CategoryId, t.Amount, t.AccountId, t.Description)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) DeleteTransaction(t *models.DeleteTransaction) error {
	_, err := d.db.Exec(deleteTransactionSchema, t.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateTransaction(t *models.Transaction) error {
	_, err := d.db.Exec(updateTransactionSchema, t.UserId, t.CategoryId, t.Amount, t.AccountId, t.Description, t.Id)
	if err != nil {
		return err
	}
	return err
}
