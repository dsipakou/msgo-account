package server

import (
	"fmt"
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func (a *Api) GetAccountsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accounts, err := a.DB.GetAccounts()
		if err != nil {
			log.Printf("Cannot get accounts, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonAccountResponse, len(accounts))
		for idx, account := range accounts {
			resp[idx] = utils.MapAccountToJson(account)
		}

		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) DeleteAccountHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start delete account...")
		request := models.JsonAccountDelete{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		t := &models.JsonAccountDelete{
			Id: request.Id,
		}

		err = a.DB.DeleteAccount(t)
		if err != nil {
			log.Printf("Cannot delete account. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}

func (a *Api) CreateAccountHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonAccountCreate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

    account, err := a.DB.CreateAccount(&request)
		if err != nil {
			log.Printf("Cannot save account in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapAccountToJson(account)
		utils.SendResponse(w, r, resp, http.StatusCreated)
	}
}

func (a *Api) UpdateAccountHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonAccountUpdate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		account := &models.Account{
			Id:          request.Id,
			UserId:      request.UserId,
			Source:      request.Source,
			Amount:      request.Amount,
			Description: request.Description,
		}

		err = a.DB.UpdateAccount(account)
		if err != nil {
			log.Printf("Cannot update account. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}
