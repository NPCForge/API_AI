package sharedModel

type StatusRequest struct {
	Action string `json:"string"`
	Token  string `json:"token"`
}

type StatusResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
