package server

import (
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func (a *Api) GetBudgetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		budget, err := a.DB.GetBudget()
		if err != nil {
			log.Printf("Cannot get budget, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonBudgetResponse, len(budget))
		for idx, item := range budget {
			resp[idx] = utils.MapBudgetToJson(item)
		}
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) CreateBudgetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonBudgetCreate{}
    log.Println(r)
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		budget, err := a.DB.CreateBudget(&request)
		if err != nil {
			log.Printf("Cannot save budget in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapBudgetToJson(budget)
		utils.SendResponse(w, r, resp, http.StatusCreated)
	}
}

func (a *Api) DeleteBudgetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &models.JsonBudgetDelete{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		err = a.DB.DeleteBudget(request)
		if err != nil {
			log.Printf("Cannot delete budget. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		utils.SendResponse(w, r, nil, http.StatusNoContent)
	}
}

func (a *Api) UpdateBudgetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonBudgetUpdate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		budget, err := a.DB.UpdateBudget(&request)
		if err != nil {
			log.Printf("Cannot update budget. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapBudgetToJson(budget)
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}
