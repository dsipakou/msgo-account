package server

import (
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func (a *Api) GetTransactionsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactions, err := a.DB.GetTransactions()
		if err != nil {
			log.Printf("Cannot get transactions, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonTransactionResponse, len(transactions))
		for idx, transaction := range transactions {
			resp[idx] = utils.MapTransactionToJson(transaction)
		}

		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) CreateTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonTransactionCreate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		t := &models.Transaction{
			Id:              0,
			UserId:          request.UserId,
			CategoryId:      request.CategoryId,
			AccountId:       request.AccountId,
			Amount:          request.Amount,
			TransactionDate: request.TransactionDate,
			Description:     request.Description,
		}

		id, err := a.DB.CreateTransaction(t)
		if err != nil {
			log.Printf("Cannot save transaction in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		t.Id = int32(id)

		resp := utils.MapTransactionToJson(t)
		utils.SendResponse(w, r, resp, http.StatusCreated)
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

		t := &models.JsonTransactionDelete{
			Id: request.Id,
		}
		err = a.DB.DeleteTransaction(t)
		if err != nil {
			log.Printf("Cannot delete transaction. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		utils.SendResponse(w, r, nil, http.StatusNoContent)
	}
}

func (a *Api) UpdateTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonTransactionUpdate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		t := &models.Transaction{
			Id:              request.Id,
			UserId:          request.UserId,
			CategoryId:      request.CategoryId,
			AccountId:       request.AccountId,
			Amount:          request.Amount,
			TransactionDate: request.TransactionDate,
			Description:     request.Description,
		}

		err = a.DB.UpdateTransaction(t)
		if err != nil {
			log.Printf("Cannot update transaction. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		utils.SendResponse(w, r, nil, http.StatusOK)
	}
}
