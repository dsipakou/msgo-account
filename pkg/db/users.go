package db

import (
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
)

type UserDB interface {
	GetUser(email string) (models.User, error)
	GetUsers() ([]*models.User, error)
	CreateUser(u *models.User) error
	DeleteUser(u *models.JsonUserDelete) error
	UpdateUser(u *models.JsonUserUpdate) error
}

func (d *DB) GetUser(email string) (models.User, error) {
	var user models.User
	err := d.db.Get(&user, getUserSchema, email)
	log.Println(err, user)
	if err != nil {
		return user, err
	}

	return user, nil
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
	password := utils.GetHash([]byte(a.Password))
	res, err := d.db.Exec(insertUserSchema, a.Name, a.Email, password)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) DeleteUser(a *models.JsonUserDelete) error {
	_, err := d.db.Exec(deleteUserSchema, a.Id)
	if err != nil {
		return err
	}

	return err
}

func (d *DB) UpdateUser(t *models.JsonUserUpdate) error {
  password := utils.GetHash([]byte(t.Password))
	_, err := d.db.Exec(updateUserSchema, t.Name, t.Email, password)
	if err != nil {
		return err
	}
	return err
}

func (d *DB) ResetUser(t *models.JsonResetUserRequest) error {
  password := utils.GetHash([]byte(t.Password))
	_, err := d.db.Exec(resetUserSchema, password, t.Email)
	if err != nil {
		return err
	}
	return err
}
