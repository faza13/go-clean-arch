package resources

type JWTResource struct {
	Iat  string `json:"iat"`
	sub  string `json:"sub"`
	Name string `json:"name"`
	id   uint64 `json:"id"`
}
