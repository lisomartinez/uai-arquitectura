package model

type User struct {
	UserId    uint64 `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
}
