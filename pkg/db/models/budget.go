package models

// DB structures
type Budget struct {
	Id          int32   `db:"id"`
	BudgetDate  string  `db:"budget_date"`
	CategoryId  *int32  `db:"category_id"`
	Title       string  `db:"title"`
	Amount      float32 `db:"amount"`
	Description string  `db:"description"`
	IsCompleted bool    `db:"is_completed"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

type ExtendedBudget struct {
	Id                      int32   `db:"id"`
	BudgetDate              string  `db:"budget_date"`
	CategoryId              *int32  `db:"category_id"`
	Title                   string  `db:"title"`
	Amount                  float32 `db:"amount"`
	Description             string  `db:"description"`
	IsCompleted             bool    `db:"is_completed"`
	CreatedAt               string  `db:"created_at"`
	UpdatedAt               string  `db:"updated_at"`
	SpentInOriginalCurrency *float32 `db:"spent_in_original_currency"`
	SpentInBaseCurrency     *float32 `db:"spent_in_base_currency"`
}

type BudgetUsage struct {
	Name   string  `db:"name"`
	Amount float32 `db:"amount"`
}

// Json structures

type JsonBudgetCreate struct {
	BudgetDate  string  `json:"budgetDate"`
	CategoryId  *int32  `json:"categoryId,omitempty"`
	Title       string  `json:"title"`
	Amount      float32 `json:"amount,omitempty"`
	Description string  `json:"description"`
}

type JsonBudgetUpdate struct {
	Id          int32   `json:"id"`
	BudgetDate  string  `json:"budgetDate"`
	CategoryId  *int32  `json:"categoryId,omitempty"`
	Title       string  `json:"title"`
	Amount      float32 `json:"amount,omitempty"`
	Description string  `json:"description"`
	IsCompleted bool    `json:"isCompleted"`
}

type JsonBudgetDelete struct {
	Id int32 `json:"id"`
}

type JsonBudgetResponse struct {
	Id          int32   `json:"id"`
	BudgetDate  string  `json:"budgetDate"`
	Title       string  `json:"title"`
	Amount      float32 `json:"amount"`
	CategoryId  *int32  `json:"categoryId"`
	Description string  `json:"description"`
	IsCompleted bool    `json:"isCompleted"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type JsonExtendedBudgetResponse struct {
	Id                      int32   `json:"id"`
	BudgetDate              string  `json:"budgetDate"`
	Title                   string  `json:"title"`
	Amount                  float32 `json:"amount"`
	CategoryId              *int32  `json:"categoryId"`
	Description             string  `json:"description"`
	IsCompleted             bool    `json:"isCompleted"`
	CreatedAt               string  `json:"createdAt"`
	UpdatedAt               string  `json:"updatedAt"`
	SpentInOriginalCurrency *float32 `json:"spentInOriginalCurrency"`
	SpentInBaseCurrency     *float32 `json:"spentInBaseCurrency"`
}

type JsonBudgetUsageResponse struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
}
