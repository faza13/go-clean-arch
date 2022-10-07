package resources

type UserStoreRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	Password           string `json:"password" form:"password" validate:"required"`
	NewPassword        string `json:"new_password" form:"new_password" validate:"required,min=6"`
	NewPasswordConfirm string `json:"new_password_confirm" form:"new_password_confirm" validate:"required,eqfield=NewPassword"`
}
