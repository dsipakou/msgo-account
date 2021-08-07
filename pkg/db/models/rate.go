package models

type Rate struct {
	Id          int32   `db:"id"`
	CurrencyId  int32   `db:"currency_id"`
	RateDate    string  `db:"rate_date"`
	Rate        float32 `db:"rate"`
	Description string  `db:"description"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

type JsonRateCreate struct {
	CurrencyId  int32   `json:"currencyId"`
	RateDate    string  `json:"rateDate"`
	Rate        float32 `json:"rate"`
	Description string  `json:"description"`
}

type JsonRateUpdate struct {
	Id          int32   `json:"id"`
	CurrencyId  int32   `json:"currencyId"`
	RateDate    string  `json:"rateDate"`
	Rate        float32 `json:"rate"`
	Description string  `json:"description"`
}

type JsonRateDelete struct {
	Id int32 `json:"id"`
}

type JsonRateResponse struct {
	Id          int32   `json:"id"`
	CurrencyId  int32   `json:"currencyId"`
	RateDate    string  `json:"rateDate"`
	Rate        float32 `json:"rate"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}
