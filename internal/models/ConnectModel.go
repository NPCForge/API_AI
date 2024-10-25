package models

// ConnectRequest représente la structure des données attendues dans la requête
type ConnectRequest struct {
	Token string `json:"BearerToken"`
}

// ConnectResponse représente la structure de la réponse à renvoyer
type ConnectResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
