package model

import (
	"time"
)

type CreateOrderRequest struct {
	OrderID      uint64   `json:"order_id"`
	UserID       uint64   `json:"user_id"`
	RestaurantID uint64   `json:"restaurant_id"`
	Items        []uint64 `json:"items"`
}

type Order struct {
	ID           string     `bson:"_id"`
	OrderID      uint64     `json:"order_id" bson:"order_id"`
	UserID       uint64     `json:"user_id" bson:"user_id"`
	RestaurantID uint64     `json:"restaurant_id" bson:"restaurant_id"`
	Items        []MenuItem `json:"items" bson:"items"`
	Total        float64    	`json:"total" bson:"total"`
	DateCreated  time.Time  `json:"date_created" bson:"date_created"`
	LastUpdated  time.Time  `json:"last_updated" bson:"last_updated"`
	Version      uint64     `json:"version" bson:"version"`
}
