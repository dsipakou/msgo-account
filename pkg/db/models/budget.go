package models

// DB structures
type Budget struct {
	Id          int32   `db:"id"`
	BudgetDate  string  `db:"budget_date"`
	Title       string  `db:"title"`
	Amount      float32 `db:"amount"`
	Description string  `db:"description"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

// Json structures

type JsonBudgetCreate struct {
	BudgetDate  string  `json:"budgetDate"`
	Title       string  `json:"title"`
	Amount      float32 `json:"amount,string,omitempty"`
	Description string  `json:"description"`
}

type JsonBudgetUpdate struct {
	Id          int32   `json:"id"`
	BudgetDate  string  `json:"budgetDate"`
	Title       string  `json:"title"`
	Amount      float32 `json:"amount,string,omitempty"`
	Description string  `json:"description"`
}

type JsonBudgetDelete struct {
	Id int32 `json:"id"`
}

type JsonBudgetResponse struct {
	Id          int32   `json:"id"`
	BudgetDate  string  `json:"budgetDate"`
	Title       string  `json:"title"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}
