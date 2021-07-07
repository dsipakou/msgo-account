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

func (a *Api) CreateTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.TransactionRequest{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		t := &models.Transaction{
			Id:       0,
			UserId:   request.UserId,
			Category: request.Category,
			Amount:   request.Amount,
		}

		err = a.DB.CreateTransaction(t)
		if err != nil {
			log.Printf("Cannot save post in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapTransactionToJson(t)
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
