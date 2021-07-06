package server

import (
	"fmt"
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func (a *Api) IndexHandler() http.HandlerFunc {
  log.Println("Index requested")
	return func(w http.ResponseWriter, r *http.Request) {
    log.Println("Index called")
		fmt.Fprintf(w, "Welcome to Account API")
	}
}

func (a *Api) GetTransactionsHandler() http.HandlerFunc {
  log.Println("Get transactions requested")
	return func(w http.ResponseWriter, r *http.Request) {
    log.Println("Get transactions called")
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
