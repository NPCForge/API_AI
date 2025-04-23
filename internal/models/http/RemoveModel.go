package httpModels

type RemoveRequestRefacto struct {
	Token                string `json:"token"`
	DeleteUserIdentifier string `json:"deleteUser"`
}

type RemoveResponseRefacto struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
