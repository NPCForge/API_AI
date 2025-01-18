package websocketModels

type ConnectRequest struct {
	Action string `json:"action"`
	Token  string `json:"checksum"`
}
