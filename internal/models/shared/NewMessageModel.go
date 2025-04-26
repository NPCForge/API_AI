package sharedModel

type NewMessageRequest struct {
	Action    string   `json:"action"`
	Sender    string   `json:"sender"`
	Receivers []string `json:"receivers"`
	Message   string   `json:"message"`
	Token     string   `json:"token"`
}

type NewMessageResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
