package sharedModel

type ChatGptSimpleRequestBodyMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGptSimpleRequestBody struct {
	Messages []ChatGptSimpleRequestBodyMessage `json:"messages"`
	Model    string                            `json:"model"`
}

type ChatGPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type Message struct {
	SenderName    string
	ReceiverNames []string
	Message       string
}
