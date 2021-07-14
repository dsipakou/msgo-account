package handlers

import (
	"log"
	"msgo-account/pkg/db"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func GetTransactionsHandler(db db.GeneralDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactions, err := db.GetTransactions()
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

func CreateTransactionHandler(db db.GeneralDB) http.HandlerFunc {
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

		err = db.CreateTransaction(t)
		if err != nil {
			log.Printf("Cannot save transaction in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapTransactionToJson(t)
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func DeleteTransactionHandler(db db.GeneralDB) http.HandlerFunc {
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

		err = db.DeleteTransaction(t)
		if err != nil {
			log.Printf("Cannot delete transaction. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}

func UpdateTransactionHandler(db db.GeneralDB) http.HandlerFunc {
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

		err = db.UpdateTransaction(t)
		if err != nil {
			log.Printf("Cannot update transaction. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}
