package models

type DisconnectRequest struct {
	Action string `json:"action"`
	Token  string `json:"token"`
}
