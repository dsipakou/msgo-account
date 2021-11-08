package server

import (
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
	"time"
)

func (a *Api) GetBudgetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dateFrom := r.FormValue("dateFrom")
		dateTo := r.FormValue("dateTo")
		log.Println(dateFrom, dateTo)
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

func (a *Api) GetBudgetForPeriodHandler() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    dateFrom := r.FormValue("dateFrom")
    dateTo := r.FormValue("dateTo")
		budgetList, err := a.DB.GetBudgetForPeriod(dateFrom, dateTo)
		if err != nil {
			log.Printf("Cannot get budget usage, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonBudgetResponse, len(budgetList))
		for idx, item := range budgetList {
			resp[idx] = utils.MapBudgetToJson(item)
		}
		utils.SendResponse(w, r, resp, http.StatusOK)
  }
}

func (a *Api) GetBudgetUsageForPeriodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dateFrom, err := time.Parse("2006-01-02", r.FormValue("dateFrom"))
    if err != nil {
      log.Printf("dateFrom is not a date, input - %v \n", r.FormValue("dateFrom"))
			utils.SendResponse(w, r, "Incorrect date format", http.StatusBadRequest)
			return
    }
		dateTo, err := time.Parse("2006-01-02", r.FormValue("dateTo"))
    if err != nil {
      log.Printf("dateFrom is not a date, input - %v \n", r.FormValue("dateFrom"))
			utils.SendResponse(w, r, "Incorrect date format", http.StatusBadRequest)
			return
    }
		usage, err := a.DB.GetBudgetUsage(dateFrom.Format("2006-01-02"), dateTo.Format("2006-01-02"))
		if err != nil {
			log.Printf("Cannot get budget usage, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonBudgetUsageResponse, len(usage))
		for idx, item := range usage {
			resp[idx] = utils.MapBudgetUsageToJson(item)
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
