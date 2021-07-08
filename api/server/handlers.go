package server

import (
	"fmt"
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func (a *Api) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Account API")
	}
}

func (a *Api) GetTransactionsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactions, err := a.DB.GetTransactions()
		if err != nil {
			log.Printf("Cannot get transactions, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonTransaction, len(transactions))
		for idx, transaction := range transactions {
			resp[idx] = utils.MapTransactionToJson(transaction)
		}

		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) GetAccountsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accounts, err := a.DB.GetAccounts()
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

func (a *Api) CreateTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonTransactionRequest{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		t := &models.Transaction{
			Id:          0,
			UserId:      request.UserId,
			Category:    request.Category,
			AccountId:   request.AccountId,
			Amount:      request.Amount,
			Description: request.Description,
		}

		err = a.DB.CreateTransaction(t)
		if err != nil {
			log.Printf("Cannot save transaction in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapTransactionToJson(t)
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) CreateAccountHandler() http.HandlerFunc {
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

		err = a.DB.CreateAccount(account)
		if err != nil {
			log.Printf("Cannot save account in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapAccountToJson(account)
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) DeleteTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonTransactionDelete{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		t := &models.DeleteTransaction{
			Id: request.Id,
		}

		err = a.DB.DeleteTransaction(t)
		if err != nil {
			log.Printf("Cannot delete transaction. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}

func (a *Api) DeleteAccountHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

func (a *Api) UpdateTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonTransaction{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		t := &models.Transaction{
			Id:          request.Id,
			UserId:      request.UserId,
			Category:    request.Category,
			AccountId:   request.AccountId,
			Amount:      request.Amount,
			Description: request.Description,
		}

		err = a.DB.UpdateTransaction(t)
		if err != nil {
			log.Printf("Cannot update transaction. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}

func (a *Api) UpdateAccountHandler() http.HandlerFunc {
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

		err = a.DB.UpdateAccount(account)
		if err != nil {
			log.Printf("Cannot update account. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}
