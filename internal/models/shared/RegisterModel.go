package sharedModel

type RegisterRequest struct {
	Action     string `json:"action"`
	Token      string `json:"API_KEY"`
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Private string `json:"token"`
	Id      string `json:"id"`
}
