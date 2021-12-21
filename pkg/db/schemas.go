package db

var getAllTransactionsSchema = `SELECT * FROM transactions ORDER BY %s DESC`
var getGroupedTransactionsSchema = `
  SELECT 
    CONCAT(EXTRACT(YEAR FROM transaction_date), '-', EXTRACT(MONTH FROM transaction_date)) AS month, 
    EXTRACT(DAY FROM transaction_date) AS day, 
    SUM(amount) AS grouped_amount 
  FROM transactions 
  GROUP BY transaction_date, type 
  HAVING type='outcome' AND transaction_date >= '%s' AND transaction_date <= '%s' 
  ORDER BY transaction_date`
var getRangedTransactionsSchema = `
  SELECT * 
  FROM transactions
  WHERE type='outcome' AND transaction_date >= '%s' AND transaction_date <= '%s'
  ORDER BY transaction_date`
var getTransactionSchema = `SELECT * FROM transactions WHERE id=$1`
var insertTransactionSchema = `INSERT INTO transactions("user_id", "category_id", "amount", "account_id", "dest_account_id", "budget_id", "transaction_date", "type", "description") VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, created_at, updated_at`
var deleteTransactionSchema = `DELETE FROM transactions WHERE id=$1`
var updateTransactionSchema = `UPDATE transactions SET user_id=$1, category_id=$2, amount=$3, account_id=$4, dest_account_id=$5, budget_id=$6, transaction_date=$7, type=$8, description=$9, updated_at=NOW() WHERE id=$10`

var getAccountsSchema = `SELECT * FROM accounts`
var getAccountSchema = `SELECT * FROM accounts WHERE id=$1`
var insertAccountSchema = `INSERT INTO accounts(user_id, source, amount, description, is_main) VALUES($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
var deleteAccountSchema = `DELETE FROM accounts WHERE id=$1`
var updateAccountSchema = `UPDATE accounts SET user_id=$1, source=$2, amount=$3, description=$4, is_main=$5 WHERE id=$6`

var getUserSchema = `SELECT * FROM users WHERE email=$1`
var getUsersSchema = `SELECT * FROM users`
var insertUserSchema = `INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id`
var deleteUserSchema = `DELETE FROM users WHERE id=$1`
var updateUserSchema = `UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4`

var getCategoriesSchema = `SELECT * FROM categories ORDER BY name`
var getCategorySchema = `SELECT * FROM categories WHERE id=$1`
var insertCategorySchema = `INSERT INTO categories(name, parent, is_parent) VALUES($1, $2, $3) RETURNING id, created_at, updated_at`
var deleteCategorySchema = `DELETE FROM categories WHERE id=$1`
var updateCategorySchema = `UPDATE categories SET name=$1, parent=$2, is_parent=$3 WHERE id=$4`

var getAllCurrenciesSchema = `SELECT * FROM currencies`
var getCurrencySchema = `SELECT * FROM currencies WHERE id=$1`
var getDefaultCurrencySchema = `SELECT * from currencies WHERE is_default=true`
var insertCurrencySchema = `INSERT INTO currencies(code, sign, verbal_name, is_default, comments) VALUES($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
var deleteCurrencySchema = `DELETE FROM currencies WHERE id=$1`
var updateCurrencySchema = `UPDATE currencies SET code=$1, sign=$2, verbal_name=$3, is_default=$4, comments=$5 WHERE id=$6`

var getAllRatesSchema = `SELECT * FROM rates ORDER BY rate_date DESC`
var getRateSchema = `SELECT * FROM rates WHERE id=$1`
var insertRateSchema = `INSERT INTO rates(currency_id, rate_date, rate, description) VALUES($1, $2, $3, $4) RETURNING id, created_at, updated_at`
var deleteRateSchema = `DELETE FROM rates WHERE id=$1`
var updateRateSchema = `UPDATE rates SET currency_id=$1, rate_date=$2, rate=$3, description=$4 WHERE id=$5`

var getAllBudgetSchema = `SELECT * FROM budget`
var getBudgetSchema = `SELECT * FROM budget WHERE id=$1`
var getBudgetForPeriod = `SELECT * FROM budget WHERE budget_date BETWEEN '%s' AND '%s' ORDER BY title`
var insertBudgetSchema = `INSERT INTO budget(budget_date, title, amount, category_id, description) VALUES($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
var deleteBudgetSchema = `DELETE FROM budget WHERE id=$1`
var updateBudgetSchema = `UPDATE budget SET budget_date=$1, title=$2, amount=$3, category_id=$4, description=$5, is_completed=$6 WHERE id=$7`
var getPeriodBudgetUsage = `
  SELECT sum(t.amount) AS amount, c.parent AS name
  FROM transactions as t
  INNER JOIN categories c on c.id = t.category_id
  WHERE t.transaction_date >= '%s' AND t.transaction_date < '%s'
    AND t.budget_id IS NOT NULL
  GROUP BY c.parent;
`
