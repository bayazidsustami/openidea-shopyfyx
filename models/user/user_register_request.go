package user_model

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
