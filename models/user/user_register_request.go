package models

type UserRegisterRequest struct {
	Username string
	Password string
	Name     string
}

type UserLoginRequest struct {
	Username string
	Password string
}
