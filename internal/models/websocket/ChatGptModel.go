package websocketModels

type TalkToResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Message struct {
	SenderName    string
	ReceiverNames []string
	Message       string
}
