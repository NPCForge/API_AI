package websocketModels

type ConnectRequestRefacto struct {
	Password   string `json:"password"`
	Identifier string `json:"identifier"`
	Action     string `json:"action"`
}

type ConnectRequest struct {
	Action string `json:"action"`
	Token  string `json:"checksum"`
}
