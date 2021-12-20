package models

type Currency struct {
	Id         int32  `db:"id"`
	Code       string `db:"code"`
	Sign       string `db:"sign"`
	VerbalName string `db:"verbal_name"`
	IsDefault  bool   `db:"is_default"`
	IsBase     bool   `db:"is_base"`
	Comments   string `db:"comments"`
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
}

type JsonCurrencyCreate struct {
	Code       string `json:"code"`
	Sign       string `json:"sign"`
	VerbalName string `json:"verbalName"`
	IsDefault  bool   `json:"isDefault"`
	IsBase     bool   `json:"isBase"`
	Comments   string `json:"comments"`
}

type JsonCurrencyUpdate struct {
	Id         int32  `json:"id"`
	Code       string `json:"code"`
	Sign       string `json:"sign"`
	VerbalName string `json:"verbalName"`
	IsDefault  bool   `json:"isDefault"`
	IsBase     bool   `json:"isBase"`
	Comments   string `json:"comments"`
}

type JsonCurrencyDelete struct {
	Id int32 `json:"id"`
}

type JsonCurrencyResponse struct {
	Id         int32  `json:"id"`
	Code       string `json:"code"`
	Sign       string `json:"sign"`
	VerbalName string `json:"verbalName"`
	IsDefault  bool   `json:"isDefault"`
	IsBase     bool   `json:"isBase"`
	Comments   string `json:"comments"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}
