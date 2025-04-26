package sharedModel

type RemoveUserRequest struct {
	Action               string `json:"string"`
	Token                string `json:"token"`
	DeleteUserIdentifier string `json:"deleteUser"`
}

type RemoveUserResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
