package websocketModels

type RemoveRequest struct {
	Action string `json:"action"`
	Token  string `json:"token"`
}
