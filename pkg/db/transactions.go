package db

import (
	"log"
	"msgo-account/pkg/db/models"
)

type TransactionDB interface {
	GetTransactions() ([]models.Transaction, error)
	CreateTransaction(m *models.JsonTransactionCreate) (models.Transaction, error)
	DeleteTransaction(m *models.JsonTransactionDelete) error
	UpdateTransaction(m *models.JsonTransactionUpdate) (models.Transaction, error)
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

  ratedAmount := m.Amount * m.Rate

	err = stmt.QueryRow(
		m.UserId,
		m.CategoryId,
    ratedAmount,
		m.AccountId,
		m.TransactionDate,
		m.Type,
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
		Amount:          ratedAmount,
		TransactionDate: m.TransactionDate,
		Type:            m.Type,
		Description:     m.Description,
		CreatedAt:       created_at,
		UpdatedAt:       updated_at,
	}

	return transaction, err
}

func (d *DB) DeleteTransaction(m *models.JsonTransactionDelete) error {
	_, err := d.db.Exec(deleteTransactionSchema, m.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateTransaction(m *models.JsonTransactionUpdate) (models.Transaction, error) {
  ratedAmount := m.Amount * m.Rate

	_, err := d.db.Exec(
		updateTransactionSchema,
		m.UserId,
		m.CategoryId,
    ratedAmount,
		m.AccountId,
		m.TransactionDate,
		m.Type,
		m.Description,
		m.Id,
	)

	if err != nil {
		return models.Transaction{}, err
	}

	var transaction models.Transaction
	err = d.db.Get(&transaction, getTransactionSchema, m.Id)

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}
