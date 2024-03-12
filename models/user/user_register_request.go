package user_model

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,min=5,max=15"`
	Password string `json:"password" validate:"required,min=5,max=15"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
}
