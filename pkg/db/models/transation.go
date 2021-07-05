package models

type Transaction struct {
	Id       int32  `db:"id"`
	UserId   int64  `db:"user_id"`
	Category string `db:"category"`
	Amount   int64  `db:"amount"`
}

type TransactionRequest struct {
	UserId   int64  `db:"user_id"`
	Category string `db:"category"`
	Amount   int64  `db:"amount"`
}

type JsonTransaction struct {
	Id       int32  `json:"id"`
	UserId   int64  `json:"user_id"`
	Category string `json:"category"`
	Amount   int64  `json:"amount"`
}
