package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Gender   int    `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
