package utils

import (
	"encoding/json"
	"log"
	"msgo-account/pkg/db/models"
	"net/http"
)

func Parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func MapTransactionToJson(m models.Transaction) models.JsonTransactionResponse {
	return models.JsonTransactionResponse{
		Id:              m.Id,
		UserId:          m.UserId,
		CategoryId:      m.CategoryId,
		Amount:          m.Amount,
		AccountId:       m.AccountId,
		TransactionDate: m.TransactionDate,
		Description:     m.Description,
    Type: m.Type,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func MapAccountToJson(m models.Account) models.JsonAccountResponse {
	return models.JsonAccountResponse{
		Id:          m.Id,
		UserId:      m.UserId,
		Source:      m.Source,
		Amount:      m.Amount,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func MapUserToJson(a *models.User) models.JsonUser {
	return models.JsonUser{
		Id:        a.Id,
		Name:      a.Name,
		Email:     a.Email,
		Password:  a.Password,
		CreatedAt: a.CreatedAt,
	}
}

func MapCategoryToJson(m models.Category) models.JsonCategoryResponse {
	return models.JsonCategoryResponse{
		Id:        m.Id,
		Name:      m.Name,
		Parent:    m.Parent,
		IsParent:  m.IsParent,
		IsSystem:  m.IsSystem,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func SendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		log.Printf("Cannot format json, err=%v\n", err)
	}
}
