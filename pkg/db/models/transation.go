package models

type Transaction struct {
	Id          int32  `db:"id"`
	UserId      int64  `db:"user_id"`
	Category    string `db:"category"`
	AccountId   int32  `db:"account_id"`
	Amount      int64  `db:"amount"`
	Description string `db:"description"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

type TransactionRequest struct {
	UserId      int64  `json:"userId"`
	Category    string `json:"category"`
	Amount      int64  `json:"amount"`
	AccountId   int32  `json:"account_id"`
	Description string `json:"description"`
}

type JsonTransactionDelete struct {
	Id int32 `json:id`
}

type DeleteTransaction struct {
	Id int32 `db:id`
}

type JsonTransaction struct {
	Id          int32  `json:"id"`
	UserId      int64  `json:"userId"`
	Category    string `json:"category"`
	AccountId   int32  `json:"accountId"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
