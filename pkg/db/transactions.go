package db

import (
	"log"
	"msgo-account/pkg/db/models"
)

type TransactionDB interface {
	GetTransactions() ([]models.Transaction, error)
	CreateTransaction(m *models.JsonTransactionCreate) (models.Transaction, error)
	DeleteTransaction(t *models.JsonTransactionDelete) error
	UpdateTransaction(t *models.JsonTransactionUpdate) (models.Transaction, error)
}

func (d *DB) GetTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := d.db.Select(&transactions, getAllTransactionsSchema)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (d *DB) CreateTransaction(m *models.JsonTransactionCreate) (models.Transaction, error) {
	stmt, err := d.db.Prepare(insertTransactionSchema)
	if err != nil {
		log.Fatal(err)
		return models.Transaction{}, err
	}
	defer stmt.Close()

	var id int
	var created_at string
	var updated_at string

	err = stmt.QueryRow(
		m.UserId,
		m.CategoryId,
		m.Amount,
		m.AccountId,
		m.TransactionDate,
		m.Description,
	).Scan(&id, &created_at, &updated_at)

	if err != nil {
		log.Fatal(err)
		return models.Transaction{}, err
	}

	transaction := models.Transaction{
		Id:              int32(id),
		UserId:          m.UserId,
		CategoryId:      m.CategoryId,
		AccountId:       m.AccountId,
		Amount:          m.Amount,
		TransactionDate: m.TransactionDate,
		Description:     m.Description,
		CreatedAt:       created_at,
		UpdatedAt:       updated_at,
	}

	return transaction, err
}

func (d *DB) DeleteTransaction(t *models.JsonTransactionDelete) error {
	_, err := d.db.Exec(deleteTransactionSchema, t.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateTransaction(t *models.JsonTransactionUpdate) (models.Transaction, error) {
	_, err := d.db.Exec(updateTransactionSchema, t.UserId, t.CategoryId, t.Amount, t.AccountId, t.TransactionDate, t.Description, t.Id)
	if err != nil {
		return models.Transaction{}, err
	}

	var transaction models.Transaction
	err = d.db.Get(&transaction, getTransactionSchema, t.Id)

	if err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}
