package types

type SignupRequest struct {
	Username        string `form:"username"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm-password"`
	UserType        string `form:"user-type"`
}
