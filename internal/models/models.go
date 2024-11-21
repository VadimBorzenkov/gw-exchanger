package models

import "time"

type ExchangeRate struct {
	CurrencyCode string    `json:"currency_code"` // Код валюты, например, USD
	Rate         float64   `json:"rate"`          // Курс валюты
	UpdatedAt    time.Time `json:"updated_at"`    // Время последнего обновления
}
