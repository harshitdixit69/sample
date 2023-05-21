package models

import "time"

type Product struct {
	ID           string  `json:"id"`
	Availability bool    `json:"availability"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
}

type Order struct {
	ID           string    `json:"id"`
	ProductID    string    `json:"product_id"`
	Quantity     int       `json:"quantity"`
	OrderValue   float64   `json:"order_value"`
	DispatchDate time.Time `json:"dispatch_date,omitempty"`
	OrderStatus  string    `json:"order_status"`
	Product      *Product  `json:"-"`
}

type UpdateData struct {
	OrderStatus  string    `json:"order_status"`
	DispatchDate time.Time `json:"dispatch_date"`
}
