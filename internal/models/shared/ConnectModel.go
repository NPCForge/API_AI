package sharedModel

type ConnectRequest struct {
	Password   string `json:"password"`
	Identifier string `json:"identifier"`
	Action     string `json:"action"`
}

type ConnectResponse struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Id       string `json:"id"`
	TmpToken string `json:"token"`
}
