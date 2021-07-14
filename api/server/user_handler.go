package server

import (
	"fmt"
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func (a *Api) GetUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := a.DB.GetUsers()
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

func (a *Api) CreateUserHandler() http.HandlerFunc {
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
		err = a.DB.CreateUser(user)
		if err != nil {
			log.Printf("Cannot save user in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapUserToJson(user)
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) DeleteUserHandler() http.HandlerFunc {
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

		err = a.DB.DeleteUser(user)
		if err != nil {
			log.Printf("Cannot delete user. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}

func (a *Api) UpdateUserHandler() http.HandlerFunc {
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

		err = a.DB.UpdateUser(user)
		if err != nil {
			log.Printf("Cannot update user. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}