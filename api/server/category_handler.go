package server

import (
	"log"
	"msgo-account/pkg/db/models"
	"msgo-account/pkg/utils"
	"net/http"
)

func (a *Api) GetCategoriesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := a.DB.GetCategories()
		if err != nil {
			log.Printf("Cannot get categories, err %v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonCategoryResponse, len(categories))
		for idx, category := range categories {
			resp[idx] = utils.MapCategoryToJson(category)
		}

		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) CreateCategoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonCategoryCreate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		category, err := a.DB.CreateCategory(&request)
		if err != nil {
			log.Printf("Cannot save category in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapCategoryToJson(category)
		utils.SendResponse(w, r, resp, http.StatusCreated)
	}
}

func (a *Api) DeleteCategoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &models.JsonCategoryDelete{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		err = a.DB.DeleteCategory(request)
		if err != nil {
			log.Printf("Cannot delete category. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		utils.SendResponse(w, r, nil, http.StatusNoContent)
	}
}

func (a *Api) UpdateCategoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.JsonCategoryUpdate{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		category, err := a.DB.UpdateCategory(&request)
		if err != nil {
			log.Printf("Cannot update category. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := utils.MapCategoryToJson(category)
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}
