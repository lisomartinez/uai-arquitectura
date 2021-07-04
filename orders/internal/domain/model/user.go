package model

type User struct {
	UserID    uint64 `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
}
