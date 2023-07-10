package auth

type RegisterInput struct {
	Fullname        string `json:"full_name" binding:"required"`
	Username        string `json:"user_name" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}
