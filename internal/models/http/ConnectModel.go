package httpModels

type ConnectRequestRefacto struct {
	Password   string `json:"password"`
	Identifier string `json:"identifier"`
}

// ConnectRequest représente la structure des données attendues dans la requête
type ConnectRequest struct {
	Token string `json:"checksum"`
}

// ConnectResponse représente la structure de la réponse à renvoyer
type ConnectResponse struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	TmpToken string `json:"token"`
}
