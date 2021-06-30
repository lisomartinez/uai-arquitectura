package model

type Menu struct {
	RestaurantId uint64     `json:"restaurant_id"`
	Items        []MenuItem `json:"items"`
}

type MenuItem struct {
	Id uint64 `json:"id"`
}
