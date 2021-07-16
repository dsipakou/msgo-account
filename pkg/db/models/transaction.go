package models

type Transaction struct {
	Id              int32   `db:"id"`
	UserId          int32   `db:"user_id"`
	CategoryId      int32   `db:"category_id"`
	AccountId       int32   `db:"account_id"`
	Amount          float32 `db:"amount"`
	TransactionDate string  `db:"transaction_date"`
	Description     string  `db:"description"`
	CreatedAt       string  `db:"created_at"`
	UpdatedAt       string  `db:"updated_at"`
}

type JsonTransactionCreate struct {
	UserId          int32   `json:"userId"`
	CategoryId      int32   `json:"categoryId"`
	Amount          float32 `json:"amount,string,omitempty"`
	AccountId       int32   `json:"accountId"`
	TransactionDate string  `json:"transaction_date"`
	Description     string  `json:"description"`
}

type JsonTransactionUpdate struct {
	Id              int32   `json:"id"`
	UserId          int32   `json:"userId"`
	CategoryId      int32   `json:"categoryId"`
	Amount          float32 `json:"amount,string,omitempty"`
	AccountId       int32   `json:"accountId"`
	TransactionDate string  `json:"transaction_date"`
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
	Amount          float32 `json:"amount"`
	TransactionDate string  `json:"transaction_date"`
	Description     string  `json:"description"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}
