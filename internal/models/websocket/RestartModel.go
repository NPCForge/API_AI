package websocketModels

type RestartRequest struct {
	Action string `json:"action"`
	Token  string `json:"token"`
}