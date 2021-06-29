package models

type Transaction struct {
	UserId   int64  `db:"user_id"`
	Category string `db:"category"`
	Amount   int64  `db:"amount"`
}
