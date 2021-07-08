package db

var insertTransactionSchema = `INSERT INTO transactions(user_id, category, amount, account_id) VALUES($1, $2, $3, $4) RETURNING id`
var getTransactionsSchema = `SELECT * FROM transactions`
var deleteTransactionSchema = `DELETE FROM transactions WHERE id=$1`
var updateTransactionSchema = `UPDATE transactions SET user_id=$1, category=$2, amount=$3, account_id=$4 WHERE id=$5`

var insertAccountSchema = `INSERT INTO accounts(user_id, source, amount, description) VALUES($1, $2, $3, $4) RETURNING id`
var getAccountsSchema = `SELECT * FROM accounts`
var deleteAccountSchema = `DELETE FROM accounts WHERE id=$1`
var updateAccountSchema = `UPDATE accounts SET user_id=$1, source=$2, amount=$3, desciption=$4 WHERE id=$5`
