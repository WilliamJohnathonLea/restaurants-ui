package types

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Check    string `form:"check"`
}
