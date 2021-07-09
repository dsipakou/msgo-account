package db

import (
	"fmt"
	"msgo-account/pkg/db/models"
)

type AccountDB interface {
	GetAccounts() ([]*models.Account, error)
	CreateAccount(a *models.Account) error
	DeleteAccount(a *models.JsonAccountDelete) error
	UpdateAccount(a *models.Account) error
}

func (d *DB) GetAccounts() ([]*models.Account, error) {
	var accounts []*models.Account
	err := d.db.Select(&accounts, getAccountsSchema)
	if err != nil {
		return accounts, err
	}

	return accounts, nil
}

func (d *DB) CreateAccount(a *models.Account) error {
  fmt.Println(a.UserId, a.Source, a.Amount, a.Description)
	res, err := d.db.Exec(insertAccountSchema, a.UserId, a.Source, a.Amount, a.Description)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) DeleteAccount(a *models.JsonAccountDelete) error {
	fmt.Println("Deleting account...")
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
