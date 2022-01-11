package models

type Transaction struct {
	Id              int32   `db:"id"`
	UserId          int32   `db:"user_id"`
	CategoryId      *int32  `db:"category_id"`
	AccountId       int32   `db:"account_id"`
	CurrencyId      *int32  `db:"currency_id"`
	DestAccountId   *int32  `db:"dest_account_id"`
	Amount          float32 `db:"amount"`
	BaseAmount      float32 `db:"base_amount"`
	TransactionDate string  `db:"transaction_date"`
	Type            string  `db:"type"`
	BudgetId        *int32  `db:"budget_id"`
	Description     string  `db:"description"`
	CreatedAt       string  `db:"created_at"`
	UpdatedAt       string  `db:"updated_at"`
}

type GroupedSum struct {
	AmountSum   float32 `db:"grouped_amount"`
	OriginalSum float32 `db:"original_amount,omitempty"`
	Month       string  `db:"month"`
	Day         int32   `db:"day"`
}

type JsonTransactionsGet struct {
	Sorting string
	Limit   string
}

type JsonTransactionCreate struct {
	UserId          int32   `json:"userId"`
	CategoryId      *int32  `json:"categoryId"`
	Amount          float32 `json:"amount,string,omitempty"`
	AccountId       int32   `json:"accountId"`
	DestAccountId   *int32  `json:"destAccountId"`
	CurrencyId      *int32  `json:"currencyId"`
	BudgetId        *int32  `json:"budgetId,omitempty"`
	TransactionDate string  `json:"transactionDate"`
	Type            string  `json:"type"`
	Description     string  `json:"description"`
}

type JsonTransactionUpdate struct {
	Id              int32   `json:"id"`
	UserId          int32   `json:"userId"`
	CategoryId      *int32  `json:"categoryId"`
	Amount          float32 `json:"amount,string,omitempty"`
	AccountId       int32   `json:"accountId"`
	DestAccountId   *int32  `json:"destAccountId"`
	CurrencyId      int32   `json:"currencyId"`
	BudgetId        *int32  `json:"budgetId,omitempty"`
	TransactionDate string  `json:"transactionDate"`
	Type            string  `json:"type"`
	Description     string  `json:"description"`
}

type JsonTransactionDelete struct {
	Id int32 `json:"id"`
}

type JsonTransactionResponse struct {
	Id              int32   `json:"id"`
	UserId          int32   `json:"userId"`
	CategoryId      *int32  `json:"categoryId"`
	AccountId       int32   `json:"accountId"`
	DestAccountId   *int32  `json:"destAccountId"`
	BudgetId        *int32  `json:"budgetId"`
	CurrencyId      *int32  `json:"currencyId"`
	Amount          float32 `json:"amount"`
	BaseAmount      float32 `json:"baseAmount"`
	TransactionDate string  `json:"transactionDate"`
	Type            string  `json:"type"`
	Description     string  `json:"description"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

type JsonTransactionsForMonthRequest struct {
	DateFrom     string `json:"dateFrom"`
	DateTo       string `json:"dateTo"`
	CurrencyCode string `json:"currency"`
}

type JsonTransactionsForMonthResponse struct {
	AmountSum float32 `json:"amountSum"`
	Month     string  `json:"month"`
	Day       int32   `json:"day"`
}
