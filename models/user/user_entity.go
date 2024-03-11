package models

type User struct {
	UserId   rune
	Username string
	Password string
	Name     string
}

type UserToken struct {
	TokenId     rune
	AccessToken string
}
