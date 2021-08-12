package db

var getAllTransactionsSchema = `SELECT * FROM transactions ORDER BY id DESC`
var getGroupedTransactionsSchema = `SELECT SUM(amount) FROM transactions GROUP BY transaction_date, type HAVING type='outcome' AND transaction_date >= '2021-08-01' AND transaction_date <= '2021-08-31' ORDER BY transaction_date`
var getTransactionSchema = `SELECT * FROM transactions WHERE id=$1`
var insertTransactionSchema = `INSERT INTO transactions("user_id", "category_id", "amount", "account_id", "transaction_date", "type", "description") VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at`
var deleteTransactionSchema = `DELETE FROM transactions WHERE id=$1`
var updateTransactionSchema = `UPDATE transactions SET user_id=$1, category_id=$2, amount=$3, account_id=$4, transaction_date=$5, type=$6, description=$7, updated_at=NOW() WHERE id=$8`

var getAccountsSchema = `SELECT * FROM accounts`
var getAccountSchema = `SELECT * FROM accounts WHERE id=$1`
var insertAccountSchema = `INSERT INTO accounts(user_id, source, amount, description) VALUES($1, $2, $3, $4) RETURNING id, created_at, updated_at`
var deleteAccountSchema = `DELETE FROM accounts WHERE id=$1`
var updateAccountSchema = `UPDATE accounts SET user_id=$1, source=$2, amount=$3, description=$4 WHERE id=$5`

var getUsersSchema = `SELECT * FROM users`
var insertUserSchema = `INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id`
var deleteUserSchema = `DELETE FROM users WHERE id=$1`
var updateUserSchema = `UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4`

var getCategoriesSchema = `SELECT * FROM categories`
var getCategorySchema = `SELECT * FROM categories WHERE id=$1`
var insertCategorySchema = `INSERT INTO categories(name, parent, is_parent) VALUES($1, $2, $3) RETURNING id, created_at, updated_at`
var deleteCategorySchema = `DELETE FROM categories WHERE id=$1`
var updateCategorySchema = `UPDATE categories SET name=$1, parent=$2, is_parent=$3 WHERE id=$4`

var getAllCurrenciesSchema = `SELECT * FROM currencies`
var getCurrencySchema = `SELECT * FROM currencies WHERE id=$1`
var insertCurrencySchema = `INSERT INTO currencies(code, sign, verbal_name, is_default, comments) VALUES($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
var deleteCurrencySchema = `DELETE FROM currencies WHERE id=$1`
var updateCurrencySchema = `UPDATE currencies SET code=$1, sign=$2, verbal_name=$3, is_default=$4, comments=%5 WHERE id=$6`

var getAllRatesSchema = `SELECT * FROM rates`
var getRateSchema = `SELECT * FROM rates WHERE id=$1`
var insertRateSchema = `INSERT INTO rates(currency_id, rate_date, rate, description) VALUES($1, $2, $3, $4) RETURNING id, created_at, updated_at`
var deleteRateSchema = `DELETE FROM rates WHERE id=$1`
var updateRateSchema = `UPDATE rates SET currency_id=$1, rate_date=$2, rate=$3, description=$4 WHERE id=$5`
