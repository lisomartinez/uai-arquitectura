package model

type Menu struct {
	RestaurantID uint64     `json:"restaurant_id"`
	Items        []MenuItem `json:"items"`
}

type MenuItem struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}
