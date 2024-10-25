package types

type LoginRequest struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
