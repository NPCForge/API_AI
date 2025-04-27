package sharedModel

type RemoveUserRequest struct {
	Action   string `json:"string"`
	Token    string `json:"token"`
	UserName string `json:"username"`
}

type RemoveUserResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
