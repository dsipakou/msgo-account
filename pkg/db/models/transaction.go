package models

import "database/sql"

type Transaction struct {
	Id              int32         `db:"id"`
	UserId          int32         `db:"user_id"`
	CategoryId      int32         `db:"category_id"`
	AccountId       int32         `db:"account_id"`
	CurrencyId      sql.NullInt32 `db:"currency_id"`
	Amount          float32       `db:"amount"`
	TransactionDate string        `db:"transaction_date"`
	Type            string        `db:"type"`
	BudgetId        *int32        `db:"budget_id"`
	Description     string        `db:"description"`
	CreatedAt       string        `db:"created_at"`
	UpdatedAt       string        `db:"updated_at"`
}

type GroupedSum struct {
	AmountSum float32 `db:"grouped_amount"`
	Month     string  `db:"month"`
	Day       int32   `db:"day"`
}

type JsonTransactionCreate struct {
	UserId          int32   `json:"userId"`
	CategoryId      int32   `json:"categoryId"`
	Amount          float32 `json:"amount,string,omitempty"`
	Rate            float32 `json:"rate"`
	AccountId       int32   `json:"accountId"`
	CurrencyId      int32   `json:"currencyId"`
	BudgetId        *int32  `json:"budgetId,omitempty"`
	TransactionDate string  `json:"transactionDate"`
	Type            string  `json:"type"`
	Description     string  `json:"description"`
}

type JsonTransactionUpdate struct {
	Id              int32   `json:"id"`
	UserId          int32   `json:"userId"`
	CategoryId      int32   `json:"categoryId"`
	Amount          float32 `json:"amount,string,omitempty"`
	Rate            float32 `json:"rate"`
	AccountId       int32   `json:"accountId"`
	CurrencyId      int32   `json:"currencyId"`
	BudgetId        *int32  `json:"budgetId,omitempty"`
	TransactionDate string  `json:"transactionDate"`
	Type            string  `json:"type"`
	Description     string  `json:"description"`
}

type JsonTransactionDelete struct {
	Id int32 `json:"id"`
}

type JsonTransactionResponse struct {
	Id              int32         `json:"id"`
	UserId          int32         `json:"userId"`
	CategoryId      int32         `json:"categoryId"`
	AccountId       int32         `json:"accountId"`
	BudgetId        *int32 `json:"budgetId"`
	Amount          float32       `json:"amount"`
	TransactionDate string        `json:"transactionDate"`
	Type            string        `json:"type"`
	Description     string        `json:"description"`
	CreatedAt       string        `json:"createdAt"`
	UpdatedAt       string        `json:"updatedAt"`
}

type JsonTransactionsForMonthRequest struct {
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}

type JsonTransactionsForMonthResponse struct {
	AmountSum float32 `json:"amountSum"`
	Month     string  `json:"month"`
	Day       int32   `json:"day"`
}
