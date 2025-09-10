package sharedModel

type MakeDecisionRequest struct {
	Action    string `json:"action"`
	Checksum  string `json:"checksum"`
	Message   string `json:"message"`
	Token     string `json:"token"`
	RequestID string `json:"requestId"`
}

type MakeDecisionResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
