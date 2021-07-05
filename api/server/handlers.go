package server

import (
	"fmt"
	"log"
	"msgo-account/pkg/db/models"
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
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonTransaction, len(transactions))
		for idx, transaction := range transactions {
			resp[idx] = mapTransactionToJson(transaction)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}
