package db

import (
	"log"
	"msgo-account/pkg/db/models"
)

type CurrencyDB interface {
	GetCurrencies() ([]models.Currency, error)
	CreateCurrency(m *models.JsonCurrencyCreate) (models.Currency, error)
	DeleteCurrency(m *models.JsonCurrencyDelete) error
	UpdateCurrency(m *models.JsonCurrencyUpdate) (models.Currency, error)
}

func (d *DB) GetCurrencies() ([]models.Currency, error) {
	var currencies []models.Currency
	err := d.db.Select(&currencies, getAllCurrenciesSchema)
	if err != nil {
		return currencies, err
	}

	return currencies, nil
}

func (d *DB) CreateCurrency(m *models.JsonCurrencyCreate) (models.Currency, error) {
	stmt, err := d.db.Prepare(insertCurrencySchema)
	if err != nil {
		log.Fatal(err)
		return models.Currency{}, err
	}
	defer stmt.Close()

	var id int
	var created_at string
	var updated_at string

	err = stmt.QueryRow(
		m.Code,
		m.Sign,
		m.VerbalName,
    m.IsDefault,
		m.Comments,
	).Scan(&id, &created_at, &updated_at)

	if err != nil {
		log.Fatal(err)
		return models.Currency{}, err
	}

	currency := models.Currency{
		Id:         int32(id),
		Code:       m.Code,
		Sign:       m.Sign,
		VerbalName: m.VerbalName,
    IsDefault: m.IsDefault,
		Comments:   m.Comments,
		CreatedAt:  created_at,
		UpdatedAt:  updated_at,
	}

	return currency, err
}

func (d *DB) DeleteCurrency(m *models.JsonCurrencyDelete) error {
	_, err := d.db.Exec(deleteCurrencySchema, m.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateCurrency(m *models.JsonCurrencyUpdate) (models.Currency, error) {
	_, err := d.db.Exec(
		updateCurrencySchema,
		m.Code,
		m.Sign,
		m.VerbalName,
    m.IsDefault,
		m.Comments,
		m.Id,
	)
	if err != nil {
		return models.Currency{}, err
	}

	var currency models.Currency
	err = d.db.Get(&currency, getCurrencySchema, m.Id)

	if err != nil {
		return models.Currency{}, err
	}

	return currency, nil
}
