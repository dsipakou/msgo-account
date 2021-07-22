package db

import (
	"fmt"
	"log"
	"msgo-account/pkg/db/models"
)

type TransactionDB interface {
	GetTransactions() ([]*models.Transaction, error)
	CreateTransaction(t *models.Transaction) (int, error)
	DeleteTransaction(t *models.JsonTransactionDelete) error
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

func (d *DB) CreateTransaction(t *models.Transaction) (int, error) {
	stmt, err := d.db.Prepare(insertTransactionSchema)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	defer stmt.Close()

	var id int
	var created_at string
	var updated_at string

	err = stmt.QueryRow(
		t.UserId,
		t.CategoryId,
		t.Amount,
		t.AccountId,
		t.TransactionDate,
		t.Description,
	).Scan(&id, &created_at, &updated_at)

	fmt.Println(id, created_at, updated_at)

	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	return id, err
}

func (d *DB) DeleteTransaction(t *models.JsonTransactionDelete) error {
	_, err := d.db.Exec(deleteTransactionSchema, t.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateTransaction(t *models.Transaction) error {
	fmt.Println(t)
	_, err := d.db.Exec(updateTransactionSchema, t.UserId, t.CategoryId, t.Amount, t.AccountId, t.TransactionDate, t.Description, t.Id)
	if err != nil {
		return err
	}
	return err
}
