package db

var insertTransactionSchema = `INSERT INTO transactions(user_id, category, amount) VALUES($1, $2, $3) RETURNING id`
var getTransactionsSchema = `SELECT * FROM transactions`
var deleteTransactionSchema = `DELETE FROM transactions WHERE id=$1`
