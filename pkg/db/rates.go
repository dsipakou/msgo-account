package db

import (
	"log"
	"msgo-account/pkg/db/models"
)

type RateDB interface {
	GetRates() ([]models.Rate, error)
	CreateRate(m *models.JsonRateCreate) (models.Rate, error)
	DeleteRate(m *models.JsonRateDelete) error
	UpdateRate(m *models.JsonRateUpdate) (models.Rate, error)
}

func (d *DB) GetRates() ([]models.Rate, error) {
	var rates []models.Rate
	err := d.db.Select(&rates, getAllRatesSchema)
	if err != nil {
		return rates, err
	}

	return rates, nil
}

func (d *DB) CreateRate(m *models.JsonRateCreate) (models.Rate, error) {
	stmt, err := d.db.Prepare(insertRateSchema)
	if err != nil {
		log.Fatal(err)
		return models.Rate{}, err
	}
	defer stmt.Close()

	var id int
	var created_at string
	var updated_at string

	err = stmt.QueryRow(
		m.CurrencyId,
		m.RateDate,
		m.Rate,
		m.Description,
	).Scan(&id, &created_at, &updated_at)

	if err != nil {
		log.Fatal(err)
		return models.Rate{}, err
	}

	rate := models.Rate{
		Id:          int32(id),
		CurrencyId:  m.CurrencyId,
		RateDate:    m.RateDate,
		Rate:        m.Rate,
		Description: m.Description,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}

	return rate, err
}

func (d *DB) DeleteRate(m *models.JsonRateDelete) error {
	_, err := d.db.Exec(deleteRateSchema, m.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateRate(m *models.JsonRateUpdate) (models.Rate, error) {
	_, err := d.db.Exec(
		updateCurrencySchema,
		m.CurrencyId,
		m.RateDate,
		m.Rate,
		m.Description,
		m.Id,
	)
	if err != nil {
		return models.Rate{}, err
	}

	var rate models.Rate
	err = d.db.Get(&rate, getRateSchema, m.Id)

	if err != nil {
		return models.Rate{}, err
	}

	return rate, nil
}
