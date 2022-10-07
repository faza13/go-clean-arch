package resources

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}
