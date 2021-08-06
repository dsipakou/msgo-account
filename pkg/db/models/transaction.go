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
	Description     string        `db:"description"`
	CreatedAt       string        `db:"created_at"`
	UpdatedAt       string        `db:"updated_at"`
}

type JsonTransactionCreate struct {
	UserId          int32   `json:"userId"`
	CategoryId      int32   `json:"categoryId"`
	Amount          float32 `json:"amount,string,omitempty"`
	AccountId       int32   `json:"accountId"`
	CurrencyId      int32   `json:"currencyId"`
	TransactionDate string  `json:"transactionDate"`
	Type            string  `json:"type"`
	Description     string  `json:"description"`
}

type JsonTransactionUpdate struct {
	Id              int32   `json:"id"`
	UserId          int32   `json:"userId"`
	CategoryId      int32   `json:"categoryId"`
	Amount          float32 `json:"amount,string,omitempty"`
	AccountId       int32   `json:"accountId"`
	CurrencyId      int32   `json:"currencyId"`
	TransactionDate string  `json:"transactionDate"`
	Type            string  `json:"type"`
	Description     string  `json:"description"`
}

type JsonTransactionDelete struct {
	Id int32 `json:"id"`
}

type JsonTransactionResponse struct {
	Id              int32   `json:"id"`
	UserId          int32   `json:"userId"`
	CategoryId      int32   `json:"categoryId"`
	AccountId       int32   `json:"accountId"`
	Currency        int32   `json:"currencyId"`
	Amount          float32 `json:"amount"`
	TransactionDate string  `json:"transactionDate"`
	Type            string  `json:"type"`
	Description     string  `json:"description"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}
