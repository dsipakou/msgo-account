package server

import (
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func (a *Api) GetCurrenciesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currencies, err := a.DB.GetCurrencies()
		if err != nil {
			log.Printf("Cannot get currencies, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonCurrencyResponse, len(currencies))
		for idx, currency := range currencies {
			resp[idx] = utils.MapCurrencyToJson(currency)
		}
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) CreateCurrencyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonCurrencyCreate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		currency, err := a.DB.CreateCurrency(&request)
		if err != nil {
			log.Printf("Cannot save currency in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapCurrencyToJson(currency)
		utils.SendResponse(w, r, resp, http.StatusCreated)
	}
}

func (a *Api) DeleteCurrencyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &models.JsonCurrencyDelete{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		err = a.DB.DeleteCurrency(request)
		if err != nil {
			log.Printf("Cannot delete currency. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		utils.SendResponse(w, r, nil, http.StatusNoContent)
	}
}

func (a *Api) UpdateCurrencyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonCurrencyUpdate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		currency, err := a.DB.UpdateCurrency(&request)
		if err != nil {
			log.Printf("Cannot update currency. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapCurrencyToJson(currency)
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}
