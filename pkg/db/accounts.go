package db

import (
	"log"
	"msgo-account/pkg/db/models"
)

type AccountDB interface {
	GetAccounts() ([]models.Account, error)
	CreateAccount(m *models.JsonAccountCreate) (models.Account, error)
	DeleteAccount(a *models.JsonAccountDelete) error
	UpdateAccount(a *models.Account) error
}

func (d *DB) GetAccounts() ([]models.Account, error) {
	var accounts []models.Account
	err := d.db.Select(&accounts, getAccountsSchema)
	if err != nil {
		return accounts, err
	}

	return accounts, nil
}

func (d *DB) CreateAccount(m *models.JsonAccountCreate) (models.Account, error) {
	stmt, err := d.db.Prepare(insertAccountSchema)
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

func (d *DB) DeleteAccount(a *models.JsonAccountDelete) error {
	_, err := d.db.Exec(deleteAccountSchema, a.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateAccount(t *models.Account) error {
	_, err := d.db.Exec(updateAccountSchema, t.UserId, t.Source, t.Amount, t.Description, t.Id)
	if err != nil {
		return err
	}
	return err
}
