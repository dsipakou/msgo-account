package handlers

import (
	"fmt"
	"log"
	"msgo-account/pkg/db"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func GetAccountsHandler(db db.GeneralDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accounts, err := db.GetAccounts()
		if err != nil {
			log.Printf("Cannot get accounts, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonAccount, len(accounts))
		for idx, account := range accounts {
			resp[idx] = utils.MapAccountToJson(account)
		}

		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func CreateAccountHandler(db db.GeneralDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonAccountGet{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		account := &models.Account{
			Id:          0,
			UserId:      request.UserId,
			Source:      request.Source,
			Amount:      request.Amount,
			Description: request.Description,
		}

		err = db.CreateAccount(account)
		if err != nil {
			log.Printf("Cannot save account in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapAccountToJson(account)
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func DeleteAccountHandler(db db.GeneralDB) http.HandlerFunc {
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

		err = db.DeleteAccount(t)
		if err != nil {
			log.Printf("Cannot delete account. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}

func UpdateAccountHandler(db db.GeneralDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonAccount{}
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

		err = db.UpdateAccount(account)
		if err != nil {
			log.Printf("Cannot update account. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}
