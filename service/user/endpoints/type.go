package endpoints

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}

type InfoRequest struct {
	ID uint64 `json:"id"`
}

type InfoResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Gender   int32  `json:"gender"`
	Phone    string `json:"phone"`
}
