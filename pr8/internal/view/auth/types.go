package auth

type NewTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
