package profile

type ChangePasswordInput struct {
	OldPassword     string `json:"old_password" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}
