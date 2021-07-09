package db

var getTransactionsSchema = `SELECT * FROM transactions`
var insertTransactionSchema = `INSERT INTO transactions(user_id, category, amount, account_id, description) VALUES($1, $2, $3, $4, $5) RETURNING id`
var deleteTransactionSchema = `DELETE FROM transactions WHERE id=$1`
var updateTransactionSchema = `UPDATE transactions SET user_id=$1, category=$2, amount=$3, account_id=$4 description=$5 WHERE id=$6`

var getAccountsSchema = `SELECT * FROM accounts`
var insertAccountSchema = `INSERT INTO accounts(user_id, source, amount, description) VALUES($1, $2, $3, $4) RETURNING id`
var deleteAccountSchema = `DELETE FROM accounts WHERE id=$1`
var updateAccountSchema = `UPDATE accounts SET user_id=$1, source=$2, amount=$3, desciption=$4 WHERE id=$5`

var getUsersSchema = `SELECT * FROM users`
var insertUserSchema = `INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id`
var deleteUserSchema = `DELETE FROM users WHERE id=$1`
var updateUserSchema = `UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4`
