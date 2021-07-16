package db

import (
	"fmt"
	"msgo-account/pkg/db/models"
)

type CategoryDB interface {
	GetCategories() ([]*models.Category, error)
	CreateCategory(t *models.JsonCategoryCreate) error
	DeleteCategory(t *models.JsonCategoryDelete) error
	UpdateCategory(t *models.JsonCategoryUpdate) error
}

func (d *DB) GetCategories() ([]*models.Category, error) {
	var categories []*models.Category
	err := d.db.Select(&categories, getCategoriesSchema)
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (d *DB) CreateCategory(t *models.JsonCategoryCreate) error {
	fmt.Println(t)
	res, err := d.db.Exec(insertCategorySchema, t.Name, t.Parent)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) DeleteCategory(t *models.JsonCategoryDelete) error {
	_, err := d.db.Exec(deleteCategorySchema, t.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateCategory(t *models.JsonCategoryUpdate) error {
	_, err := d.db.Exec(updateCategorySchema, t.Name, t.Parent, t.Id)
	if err != nil {
		return err
	}
	return err
}
