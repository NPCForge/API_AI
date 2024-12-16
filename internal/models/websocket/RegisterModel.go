package websocketModels

type RegisterRequest struct {
	Action string `json:"action"`
	Token  string `json:"checksum"`
	Name   string `json:"name"`
	Prompt string `json:"prompt"`
}
