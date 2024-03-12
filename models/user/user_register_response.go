package user_model

type UserData struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type UserResponse struct {
	Data    UserData `json:"data"`
	Message string   `json:"message"`
}
