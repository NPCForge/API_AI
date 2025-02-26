package httpModels

type MakeDecisionResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type MakeDecisionRequest struct {
	Message string `json:"message"`
}

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
