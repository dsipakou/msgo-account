package server

import (
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func (a *Api) GetRatesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rates, err := a.DB.GetRates()
		if err != nil {
			log.Printf("Cannot get rates, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonRateResponse, len(rates))
		for idx, rate := range rates {
			resp[idx] = utils.MapRateToJson(rate)
		}
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) CreateRateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonRateCreate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		rate, err := a.DB.CreateRate(&request)
		if err != nil {
			log.Printf("Cannot save rate in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapRateToJson(rate)
		utils.SendResponse(w, r, resp, http.StatusCreated)
	}
}

func (a *Api) DeleteRateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &models.JsonRateDelete{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		err = a.DB.DeleteRate(request)
		if err != nil {
			log.Printf("Cannot delete currency. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		utils.SendResponse(w, r, nil, http.StatusNoContent)
	}
}

func (a *Api) UpdateRateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonRateUpdate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		rate, err := a.DB.UpdateRate(&request)
		if err != nil {
			log.Printf("Cannot update rate. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapRateToJson(rate)
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}
