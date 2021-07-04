package model

import "time"

type CreateOrderRequest struct {
	OrderID      uint64   `json:"order_id"`
	UserID       uint64   `json:"user_id"`
	RestaurantID uint64   `json:"restaurant_id"`
	Items        []uint64 `json:"items"`
}

type Order struct {
	OrderID      uint64     `json:"order_id"`
	UserID       uint64     `json:"user_id"`
	RestaurantID uint64     `json:"restaurant_id"`
	Items        []MenuItem `json:"items"`
	Total        float64    `json:"total"`
	DateCreated  time.Time  `json:"date_created"`
	LastUpdated  time.Time  `json:"last_updated"`
	Version      uint64     `json:"version"`
}
