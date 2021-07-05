package db

import "msgo-account/pkg/db/models"

func (d *DB) CreateTransaction(t *models.Transaction) error {
	res, err := d.db.Exec(insertTransactionSchema, t.UserId, t.Category, t.Amount)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) GetTransactions() ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := d.db.Select(&transactions, getTransactionsSchema)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
