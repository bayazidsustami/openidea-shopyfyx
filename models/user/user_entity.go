package user_model

type User struct {
	UserId      int
	Username    string
	Password    string
	Name        string
	AccessToken string
}

type UserToken struct {
	TokenId     rune
	AccessToken string
}
