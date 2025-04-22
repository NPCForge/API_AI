package httpModels

type RegisterRequestRefacto struct {
	Token       string `json:"API_KEY"`
	Identifiant string `json:"identifiant"`
	Password    string `json:"password"`
}

// RegisterRequest represents the structure of the expected data in the request
type RegisterRequest struct {
	Token    string `json:"API_KEY"`
	Checksum string `json:"checksum"`
	Name     string `json:"name"`
	Prompt   string `json:"prompt"`
}

// RegisterResponse represents the structure of the response to return
type RegisterResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Private string `json:"token"`
}
