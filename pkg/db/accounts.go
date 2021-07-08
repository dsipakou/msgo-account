package db

import "msgo-account/pkg/db/models"

func (d *DB) CreateAccount(a *models.Account) error {
	res, err := d.db.Exec(insertAccountSchema, a.UserId, a.Source, a.Amount, a.Description)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}
