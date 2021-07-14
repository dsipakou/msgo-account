package handlers

import (
	"fmt"
	"log"
	"msgo-account/pkg/db"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func GetUsersHandler(db db.GeneralDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetUsers()
		if err != nil {
			log.Printf("Cannot get users, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonUser, len(users))
		for idx, user := range users {
			resp[idx] = utils.MapUserToJson(user)
		}

		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func CreateUserHandler(db db.GeneralDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonUserGet{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		user := &models.User{
			Id:       0,
			Name:     request.Name,
			Email:    request.Email,
			Password: request.Password,
		}

		fmt.Println(user)
		err = db.CreateUser(user)
		if err != nil {
			log.Printf("Cannot save user in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapUserToJson(user)
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func DeleteUserHandler(db db.GeneralDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start delete user....")
		request := models.JsonUserDelete{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		user := &models.JsonUserDelete{
			Id: request.Id,
		}

		err = db.DeleteUser(user)
		if err != nil {
			log.Printf("Cannot delete user. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}

func UpdateUserHandler(db db.GeneralDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonUser{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		user := &models.User{
			Name:     request.Name,
			Email:    request.Email,
			Password: request.Password,
		}

		err = db.UpdateUser(user)
		if err != nil {
			log.Printf("Cannot update user. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}
