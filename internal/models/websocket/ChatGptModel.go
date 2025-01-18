package websocketModels

type MakeDecisionResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type MakeDecisionRequest struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}
