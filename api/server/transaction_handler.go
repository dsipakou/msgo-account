package server

import (
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (a *Api) GetTransactionsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonTransactionsGet{}
		var transactions []models.Transaction
		switch sortParam := r.FormValue("sorting"); sortParam {
		case "date":
			request.Sorting = "transaction_date"
		case "added":
			request.Sorting = "id"
		default:
			request.Sorting = "id"
		}

    limit := r.FormValue("limit")
    if limit == "" {
      request.Limit = "15"
    } else {
      request.Limit = limit
    }

		dateFrom := r.FormValue("dateFrom")
		dateTo := r.FormValue("dateTo")

		if dateFrom != "" && dateTo != "" {
			formatedDateFrom, err := time.Parse("2006-01-02", r.FormValue("dateFrom"))
			if err != nil {
				log.Printf("dateFrom is not a date, input - %v \n", r.FormValue("dateFrom"))
				utils.SendResponse(w, r, "Incorrect date format", http.StatusBadRequest)
				return
			}

			formatedDateTo, err := time.Parse("2006-01-02", dateTo)
			if err != nil {
				log.Printf("dateTo is not a date, input - %v \n", r.FormValue("dateTo"))
				utils.SendResponse(w, r, "Incorrect date format", http.StatusBadRequest)
				return
			}

			transactions, err = a.DB.GetRangedTransactions(formatedDateFrom.Format("2006-01-02"), formatedDateTo.Format("2006-01-02"))

			if err != nil {
				log.Printf("Cannot get transactions, err %v \n", err)
				utils.SendResponse(w, r, nil, http.StatusInternalServerError)
				return
			}
    } else {
      var err error
      transactions, err = a.DB.GetTransactions(request)

      if err != nil {
        log.Printf("Cannot get transactions, err %v \n", err)
        utils.SendResponse(w, r, nil, http.StatusInternalServerError)
        return
      }
    }

		var resp = make([]models.JsonTransactionResponse, len(transactions))
		for idx, transaction := range transactions {
			resp[idx] = utils.MapTransactionToJson(transaction)
		}
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) GetGroupedTransactionsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonTransactionsForMonthRequest{}
		dateFrom := mux.Vars(r)["dateFrom"]
		dateTo := mux.Vars(r)["dateTo"]
		request.DateFrom = dateFrom
		request.DateTo = dateTo
		groupedSums, err := a.DB.GetGroupedTransactions(request)
		if err != nil {
			log.Printf("Cannot get grouped transactions, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonTransactionsForMonthResponse, len(groupedSums))
		for idx, groupSum := range groupedSums {
			resp[idx] = utils.MapGroupSumToJson(groupSum)
		}
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) GetRangedTransactionsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dateFrom := mux.Vars(r)["dateFrom"]
		dateTo := mux.Vars(r)["dateTo"]
		transactions, err := a.DB.GetRangedTransactions(dateFrom, dateTo)
		if err != nil {
			log.Printf("Cannot get ranged transactions, err %v \n", err)
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

		transaction, err := a.DB.CreateTransaction(&request)
		if err != nil {
			log.Printf("Cannot save transaction in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapTransactionToJson(transaction)
		utils.SendResponse(w, r, resp, http.StatusCreated)
	}
}

func (a *Api) DeleteTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &models.JsonTransactionDelete{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		err = a.DB.DeleteTransaction(request)
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

		transaction, err := a.DB.UpdateTransaction(&request)
		if err != nil {
			log.Printf("Cannot update transaction. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapTransactionToJson(transaction)
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}
