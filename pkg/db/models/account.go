package models

// DB structures
type Account struct {
	Id          int32   `db:"id"`
	UserId      int32   `db:"user_id"`
	Source      string  `db:"source"`
	Amount      float32 `db:"amount"`
	Description string  `db:"description"`
	IsMain      bool    `db:"is_main"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

// Json structures
type JsonAccountCreate struct {
	UserId      int32   `json:"userId"`
	Source      string  `json:"source"`
	Amount      float32 `json:"amount,string,omitempty"`
	Description string  `json:"description"`
	IsMain      bool    `json:"isMain"`
}

type JsonAccountDelete struct {
	Id int32 `db:"id"`
}

type JsonAccountUpdate struct {
	Id          int32   `json:"id"`
	UserId      int32   `json:"userId"`
	Source      string  `json:"source"`
	Amount      float32 `json:"amount,string,omitempty"`
	Description string  `json:"description"`
	IsMain      bool    `json:"isMain"`
}

type JsonAccountResponse struct {
	Id          int32   `json:"id"`
	UserId      int32   `json:"userId"`
	Source      string  `json:"source"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	IsMain      bool    `json:"isMain"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}
