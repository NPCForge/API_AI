package websocketModels

type RegisterRequest struct {
	Action   string `json:"action"`
	Token    string `json:"API_KEY"`
	Checksum string `json:"checksum"`
	Name     string `json:"name"`
	Prompt   string `json:"prompt"`
}

type RegisterRequestRefacto struct {
	Action     string `json:"action"`
	Token      string `json:"API_KEY"`
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
