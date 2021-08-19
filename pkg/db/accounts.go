package db

import (
	"log"
	"msgo-account/pkg/db/models"
)

type AccountDB interface {
	GetAccounts() ([]models.Account, error)
	CreateAccount(m *models.JsonAccountCreate) (models.Account, error)
	DeleteAccount(m *models.JsonAccountDelete) error
	UpdateAccount(m *models.JsonAccountUpdate) (models.Account, error)
}

func (d *DB) GetAccounts() ([]models.Account, error) {
	var accounts []models.Account
	err := d.db.Select(&accounts, getAccountsSchema)
	if err != nil {
		return accounts, err
	}

	return accounts, nil
}

func (d *DB) CreateAccount(m *models.JsonAccountCreate) (models.Account, error) { stmt, err := d.db.Prepare(insertAccountSchema)
	if err != nil {
		log.Fatal(err)
		return models.Account{}, err
	}
	defer stmt.Close()

	var id int
	var created_at string
	var updated_at string

	err = stmt.QueryRow(
		m.UserId,
		m.Source,
		m.Amount,
		m.Description,
	).Scan(&id, &created_at, &updated_at)

	if err != nil {
		log.Fatal(err)
		return models.Account{}, err
	}

	account := models.Account{
		Id:          int32(id),
		UserId:      m.UserId,
		Source:      m.Source,
		Amount:      m.Amount,
		Description: m.Description,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}

	return account, err
}

func (d *DB) DeleteAccount(m *models.JsonAccountDelete) error {
	_, err := d.db.Exec(deleteAccountSchema, m.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateAccount(m *models.JsonAccountUpdate) (models.Account, error) {
	_, err := d.db.Exec(updateAccountSchema, m.UserId, m.Source, m.Amount, m.Description, m.Id)
	if err != nil {
		return models.Account{}, err
	}

	var account models.Account
	err = d.db.Get(&account, getAccountSchema, m.Id)
	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}
