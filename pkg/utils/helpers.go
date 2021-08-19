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
		BudgetId:        m.BudgetId,
		TransactionDate: m.TransactionDate,
		Description:     m.Description,
		Type:            m.Type,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func MapGroupSumToJson(m models.GroupedSum) models.JsonTransactionsForMonthResponse {
	return models.JsonTransactionsForMonthResponse{
		AmountSum: m.AmountSum,
		Month:     m.Month,
		Day:       m.Day,
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

func MapBudgetToJson(m models.Budget) models.JsonBudgetResponse {
	return models.JsonBudgetResponse{
		Id:          m.Id,
		BudgetDate:  m.BudgetDate,
		Title:       m.Title,
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

func MapCurrencyToJson(m models.Currency) models.JsonCurrencyResponse {
	return models.JsonCurrencyResponse{
		Id:         m.Id,
		Code:       m.Code,
		Sign:       m.Sign,
		VerbalName: m.VerbalName,
		IsDefault:  m.IsDefault,
		Comments:   m.Comments,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func MapRateToJson(m models.Rate) models.JsonRateResponse {
	return models.JsonRateResponse{
		Id:          m.Id,
		CurrencyId:  m.CurrencyId,
		RateDate:    m.RateDate,
		Rate:        m.Rate,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
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
