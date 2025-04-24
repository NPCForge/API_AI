package websocketModels

type RemoveRequestRefacto struct {
	Action               string `json:"string"`
	Token                string `json:"token"`
	DeleteUserIdentifier string `json:"deleteUser"`
}
