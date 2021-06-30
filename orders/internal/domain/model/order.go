package model

import "time"

type CreateOrderRequest struct {
	Id           uint64   `json:"id"`
	UserId       uint64   `json:"user_id"`
	RestaurantId uint64   `json:"restaurant_id"`
	Items        []uint64 `json:"items"`
}

type Order struct {
	Id           uint64     `json:"id"`
	UserId       uint64     `json:"user_id"`
	RestaurantId uint64     `json:"restaurant_id"`
	Items        []MenuItem `json:"items"`
	Total        float64    `json:"total"`
	DateCreated  time.Time  `json:"date_created"`
	LastUpdated  time.Time  `json:"last_updated"`
	Version      uint64     `json:"version"`
}
