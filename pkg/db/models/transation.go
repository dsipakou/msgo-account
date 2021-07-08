package models

type Transaction struct {
	Id          int32   `db:"id"`
	UserId      int32   `db:"user_id"`
	Category    string  `db:"category"`
	AccountId   int32   `db:"account_id"`
	Amount      float32 `db:"amount"`
	Description string  `db:"description"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

type Account struct {
	Id          int32   `db:"id"`
	UserId      int32   `db:"user_id"`
	Source      string  `db:"source"`
	Amount      float32 `db:"amount"`
	Description string  `db:"description"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

type JsonTransactionRequest struct {
	UserId      int32   `json:"userId"`
	Category    string  `json:"category"`
	Amount      float32 `json:"amount"`
	AccountId   int32   `json:"account_id"`
	Description string  `json:"description"`
}

type JsonAccountRequest struct {
	UserId      int32   `json:"userId"`
	Source      string  `json:"source"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
}

type JsonTransactionDelete struct {
	Id int32 `json:"id"`
}

type JsonAccountDelete struct {
  Id int32 `db:"id"`
}

type DeleteTransaction struct {
	Id int32 `db:"id"`
}

type JsonTransaction struct {
	Id          int32   `json:"id"`
	UserId      int32   `json:"userId"`
	Category    string  `json:"category"`
	AccountId   int32   `json:"accountId"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type JsonAccount struct {
	Id          int32   `json:"id"`
	UserId      int32   `json:"userId"`
	Source      string  `json"source"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}
