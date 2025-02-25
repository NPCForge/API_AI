package websocketModels

type MakeDecisionResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type MakeDecisionRequest struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type TalkToResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//type TalkToRequest struct {
//	Action       string `json:"action"`
//	Message      string `json:"message"`
//	Interlocutor string `json:"interlocutor"`
//	Token        string `json:"token"`
//}
