package sharedModel

type DisconnectRequest struct {
	Action string `json:"action"`
	Token  string `json:"token"`
}

type DisconnectResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
