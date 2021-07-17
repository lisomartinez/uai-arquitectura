package model

type User struct {
	UserID    uint64 `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
