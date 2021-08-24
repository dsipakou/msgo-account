package db

import (
	"log"
	"msgo-account/pkg/db/models"
)

type BudgetDB interface {
	GetBudget() ([]models.Budget, error)
	CreateBudget(m *models.JsonBudgetCreate) (models.Budget, error)
	DeleteBudget(m *models.JsonBudgetDelete) error
	UpdateBudget(m *models.JsonBudgetUpdate) (models.Budget, error)
}

func (d *DB) GetBudget() ([]models.Budget, error) {
	var budget []models.Budget
	err := d.db.Select(&budget, getAllBudgetSchema)
	if err != nil {
		return budget, err
	}

	return budget, nil
}

func (d *DB) CreateBudget(m *models.JsonBudgetCreate) (models.Budget, error) {
	stmt, err := d.db.Prepare(insertBudgetSchema)
	if err != nil {
		log.Fatal(err)
		return models.Budget{}, err
	}
	defer stmt.Close()

	var id int
	var created_at string
	var updated_at string

	err = stmt.QueryRow(
		m.BudgetDate,
		m.Title,
		m.Amount,
		m.Description,
	).Scan(&id, &created_at, &updated_at)

	if err != nil {
		log.Fatal(err)
		return models.Budget{}, err
	}

	budget := models.Budget{
		Id:          int32(id),
		BudgetDate:  m.BudgetDate,
		Title:       m.Title,
		Amount:      m.Amount,
    IsCompleted: false,
		Description: m.Description,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}

	return budget, err
}

func (d *DB) DeleteBudget(m *models.JsonBudgetDelete) error {
	_, err := d.db.Exec(deleteBudgetSchema, m.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateBudget(m *models.JsonBudgetUpdate) (models.Budget, error) {
	_, err := d.db.Exec(updateBudgetSchema, m.BudgetDate, m.Title, m.Amount, m.Description, m.IsCompleted, m.Id)
	if err != nil {
		return models.Budget{}, err
	}

	var budget models.Budget
	err = d.db.Get(&budget, getBudgetSchema, m.Id)
	if err != nil {
		return models.Budget{}, err
	}

	return budget, nil
}
