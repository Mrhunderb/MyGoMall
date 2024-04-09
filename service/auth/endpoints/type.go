package endpoints

type GetTokenRequset struct {
	UserId uint64 `json:"user_id"`
}

type Token struct {
	Token string `json:"token"`
}

type VerifyTokenResponse struct {
	Valid bool `json:"valid"`
}
