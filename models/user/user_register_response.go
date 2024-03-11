package models

type UserData struct {
	UserId   rune
	Username string
	Name     string
}

type UserResponse struct {
	Data    UserData
	Message string
	Status  rune
}
