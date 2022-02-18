package server

import (
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (a *Api) UserLoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonUserLoginRequest{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		user, err := a.DB.GetUser(request.Email)
		if err != nil {
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		userPass := request.Password
		dbPass := user.Password
		passErr := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(userPass))
		if passErr != nil {
			utils.SendResponse(w, r, "incorrect password", http.StatusForbidden)
			return
		}

		jwtToken, err := utils.GenerateJWT(request.Email)
		if err != nil {
			utils.SendResponse(w, r, "cannot generate token", http.StatusInternalServerError)
			return
		}
		resp := &models.JsonUserToken{
			Token: jwtToken,
		}
    utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

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

func (a *Api) ResetUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonResetUserRequest{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		err = a.DB.ResetUser(&request)
		if err != nil {
			log.Printf("Cannot save user in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		if err != nil {
			log.Printf("Cannot update user. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}

func (a *Api) DeleteUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		request := models.JsonUserUpdate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		err = a.DB.UpdateUser(&request)
		if err != nil {
			log.Printf("Cannot update user. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}
