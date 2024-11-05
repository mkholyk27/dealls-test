package models

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Group    string `json:"group"`
	IsMember int    `json:"isMember"`
}

type UserDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
