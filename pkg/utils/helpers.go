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

func MapTransactionToJson(t *models.Transaction) models.JsonTransaction {
	return models.JsonTransaction{
		Id:          t.Id,
		UserId:      t.UserId,
		Category:    t.Category,
		Amount:      t.Amount,
		AccountId:   t.AccountId,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func MapAccountToJson(a *models.Account) models.JsonAccount {
	return models.JsonAccount{
		Id:          a.Id,
		UserId:      a.UserId,
		Source:      a.Source,
		Amount:      a.Amount,
		Description: a.Description,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
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
