package db

var getAllTransactionsSchema = `SELECT * FROM transactions ORDER BY %s DESC`
var getAllTransactionsExtendedSchema = `
  SELECT 
    t.*,  
    CASE WHEN t.currency_id is NULL THEN t.amount
        WHEN c.is_base THEN t.amount
        ELSE t.amount * r.rate
    END as base_amount
  FROM transactions t 
    LEFT JOIN rates r ON r.currency_id = t.currency_id AND r.rate_date = t.transaction_date
    LEFT JOIN currencies c ON c.id = t.currency_id
  ORDER BY t.%s DESC LIMIT %s;
`
var getGroupedTransactionsSchema = `
  SELECT 
    CONCAT(EXTRACT(YEAR FROM t.transaction_date), '-', EXTRACT(MONTH FROM t.transaction_date)) AS month, 
    EXTRACT(DAY FROM t.transaction_date) AS day, 
    SUM(t.amount) AS original_amount,
    SUM(CASE WHEN t.currency_id is NULL THEN t.amount
      WHEN c.is_base THEN t.amount
      ELSE t.amount * r.rate
    END) as grouped_amount 
  FROM transactions t
    LEFT JOIN rates r ON r.currency_id = t.currency_id AND r.rate_date = t.transaction_date
    LEFT JOIN currencies c ON c.id = t.currency_id
  GROUP BY t.transaction_date, t.type 
  HAVING t.type='outcome' AND t.transaction_date >= '%s' AND t.transaction_date <= '%s' 
  ORDER BY t.transaction_date;
`
var getGroupedTransactionsForCurrencySchema = `
  SELECT 
    t1.month, 
    t1.day, 
    CASE WHEN c1.is_base THEN t1.amount
      WHEN r1.rate is NULL THEN 0
      ELSE t1.amount / r1.rate 
    END as grouped_amount 
  FROM (
    SELECT 
      CONCAT(EXTRACT(YEAR FROM t.transaction_date), '-', EXTRACT(MONTH FROM t.transaction_date)) AS month, 
      EXTRACT(DAY FROM t.transaction_date) AS day, 
      SUM(CASE WHEN t.currency_id is NULL THEN t.amount
        WHEN c.is_base THEN t.amount
        ELSE t.amount * r.rate
      END) as amount,
      t.transaction_date
    FROM transactions t
      LEFT JOIN rates r ON r.currency_id = t.currency_id AND r.rate_date = t.transaction_date
      LEFT JOIN currencies c ON c.id = t.currency_id
    GROUP BY t.transaction_date, t.type, t.transaction_date
    HAVING t.type='outcome' AND t.transaction_date >= '%s' AND t.transaction_date <= '%s' 
    ORDER BY t.transaction_date
  ) as t1
    INNER JOIN currencies c1 ON c1.code = '%s'
    LEFT JOIN rates r1 ON r1.currency_id = c1.id AND r1.rate_date = t1.transaction_date
  ORDER BY r1.rate_date;
`
var getRangedTransactionsSchema = `
  SELECT 
    t.*,  
    CASE WHEN t.currency_id is NULL THEN t.amount
        WHEN c.is_base THEN t.amount
        ELSE t.amount * r.rate
    END as base_amount
  FROM transactions t
    LEFT JOIN rates r ON r.currency_id = t.currency_id AND r.rate_date = t.transaction_date
    LEFT JOIN currencies c ON c.id = t.currency_id
  WHERE type='outcome' AND transaction_date >= '%s' AND transaction_date <= '%s'
  ORDER BY transaction_date`
var getTransactionSchema = `
  SELECT 
    t.*,  
    CASE WHEN t.currency_id is NULL THEN t.amount
        WHEN c.is_base THEN t.amount
        ELSE t.amount * r.rate
    END as base_amount
  FROM transactions t 
    LEFT JOIN rates r ON r.currency_id = t.currency_id AND r.rate_date = t.transaction_date
    LEFT JOIN currencies c ON c.id = t.currency_id
  WHERE t.id=$1;
`
var insertTransactionSchema = `INSERT INTO transactions("user_id", "category_id", "amount", "account_id", "dest_account_id", "currency_id", "budget_id", "transaction_date", "type", "description") VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, created_at, updated_at`
var deleteTransactionSchema = `DELETE FROM transactions WHERE id=$1`
var updateTransactionSchema = `UPDATE transactions SET user_id=$1, category_id=$2, amount=$3, currency_id=$4, account_id=$5, dest_account_id=$6, budget_id=$7, transaction_date=$8, type=$9, description=$10, updated_at=NOW() WHERE id=$11`

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
var resetUserSchema = `UPDATE users SET password=$1 WHERE email=$2`

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
var getBudgetForPeriod = `
  SELECT b.*, t.amount as spent_in_original_currency,                                                                                                                                                             
    CASE WHEN t.currency_id is NULL THEN t.amount
      WHEN c.is_base THEN t.amount
      ELSE t.amount * r.rate
    END as spent_in_base_currency
  FROM budget AS b
    FULL JOIN transactions t on b.id = t.budget_id
    LEFT JOIN rates r ON r.currency_id = t.currency_id AND r.rate_date = t.transaction_date
    LEFT JOIN currencies c ON c.id = t.currency_id  
  WHERE b.budget_date   
  BETWEEN '%s' AND '%s'
  ORDER BY b.title
`
var insertBudgetSchema = `INSERT INTO budget(budget_date, title, amount, category_id, description) VALUES($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
var deleteBudgetSchema = `DELETE FROM budget WHERE id=$1`
var updateBudgetSchema = `UPDATE budget SET budget_date=$1, title=$2, amount=$3, category_id=$4, description=$5, is_completed=$6 WHERE id=$7`
var getBudgetPlan = `
  SELECT c.name as category_name, b.*
  FROM budget as b
  INNER JOIN categories c ON c.id = b.category_id
  WHERE b.budget_date BETWEEN '%s' AND '%s';
`
