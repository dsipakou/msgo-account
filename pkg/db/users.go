package db

import (
	"fmt"
	"msgo-account/pkg/db/models"
)

type UserDB interface {
	GetUsers() ([]*models.User, error)
	CreateUser(u *models.User) error
	DeleteUser(u *models.JsonUserDelete) error
	UpdateUser(u *models.User) error
}

func (d *DB) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := d.db.Select(&users, getUsersSchema)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (d *DB) CreateUser(a *models.User) error {
	res, err := d.db.Exec(insertUserSchema, a.Name, a.Email, a.Password)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) DeleteUser(a *models.JsonUserDelete) error {
	fmt.Println("Deleting user...")
	_, err := d.db.Exec(deleteUserSchema, a.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateUser(t *models.User) error {
	_, err := d.db.Exec(updateUserSchema, t.Name, t.Email, t.Password)
	if err != nil {
		return err
	}
	return err
}
