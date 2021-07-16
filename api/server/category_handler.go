package server

import (
	"fmt"
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

		category := &models.JsonCategoryCreate{
			Name:   request.Name,
			Parent: request.Parent,
		}

		fmt.Println(category)
		err = a.DB.CreateCategory(category)
		if err != nil {
			log.Printf("Cannot save category in DB. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		// resp := utils.MapCategoryToJson(category)
    resp := "{}"
		utils.SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *Api) DeleteCategoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start delete category....")
		request := models.JsonCategoryDelete{}
		err := utils.Parse(w, r, &request)
		if err != nil {
			log.Printf("Cannot parse body. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		category := &models.JsonCategoryDelete{
			Id: request.Id,
		}

		err = a.DB.DeleteCategory(category)
		if err != nil {
			log.Printf("Cannot delete category. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
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

		category := &models.JsonCategoryUpdate{
			Id:     request.Id,
			Name:   request.Name,
			Parent: request.Parent,
		}

		err = a.DB.UpdateCategory(category)
		if err != nil {
			log.Printf("Cannot update category. err=%v \n", err)
			utils.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
	}
}
