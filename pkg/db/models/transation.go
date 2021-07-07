package models

type Transaction struct {
	Id       int32  `db:"id"`
	UserId   int64  `db:"user_id"`
	Category string `db:"category"`
	Amount   int64  `db:"amount"`
}

type TransactionRequest struct {
	UserId   int64  `json:"userId"`
	Category string `json:"category"`
	Amount   int64  `json:"amount"`
}

type JsonTransactionDelete struct {
	Id int32 `json:id`
}

type DeleteTransaction struct {
	Id int32 `db:id`
}

type JsonTransaction struct {
	Id       int32  `json:"id"`
	UserId   int64  `json:"userId"`
	Category string `json:"category"`
	Amount   int64  `json:"amount"`
}
