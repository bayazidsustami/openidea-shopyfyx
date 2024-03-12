package user_model

type UserRegisterRequest struct {
	Username string
	Password string
	Name     string
}

type UserLoginRequest struct {
	Username string
	Password string
}
