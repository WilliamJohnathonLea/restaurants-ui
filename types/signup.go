package types

type SignupRequest struct {
	Username        string `form:"username" json:"username"`
	Password        string `form:"password" json:"password"`
	ConfirmPassword string `form:"confirm-password" json:"-"`
	UserType        string `form:"user-type" json:"type"`
}
