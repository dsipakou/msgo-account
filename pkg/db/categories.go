package db

import (
	"log"
	"msgo-account/pkg/db/models"
)

type CategoryDB interface {
	GetCategories() ([]models.Category, error)
	CreateCategory(m *models.JsonCategoryCreate) (models.Category, error)
	DeleteCategory(m *models.JsonCategoryDelete) error
	UpdateCategory(m *models.JsonCategoryUpdate) (models.Category, error)
}

func (d *DB) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := d.db.Select(&categories, getCategoriesSchema)
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (d *DB) CreateCategory(m *models.JsonCategoryCreate) (models.Category, error) {
	stmt, err := d.db.Prepare(insertCategorySchema)
	if err != nil {
		log.Fatal(err)
		return models.Category{}, err
	}
	defer stmt.Close()

	var id int
	var created_at string
	var updated_at string

	err = stmt.QueryRow(
		m.Name,
		m.Parent,
	).Scan(&id, &created_at, &updated_at)

	if err != nil {
		log.Fatal(err)
		return models.Category{}, err
	}

	category := models.Category{
		Id:        int32(id),
		Name:      m.Name,
		Parent:    m.Parent,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}

	return category, err
}

func (d *DB) DeleteCategory(m *models.JsonCategoryDelete) error {
	_, err := d.db.Exec(deleteCategorySchema, m.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateCategory(m *models.JsonCategoryUpdate) (models.Category, error) {
	_, err := d.db.Exec(updateCategorySchema, m.Name, m.Parent, m.Id)
	if err != nil {
		return models.Category{}, err
	}

	var category models.Category
	err = d.db.Get(&category, getCategorySchema, m.Id)
	if err != nil {
		return models.Category{}, err
	}

	return category, err
}
