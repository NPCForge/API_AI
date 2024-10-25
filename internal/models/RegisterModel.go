package models

// RegisterRequest représente la structure des données attendues dans la requête
type RegisterRequest struct {
	Token string `json:"BearerToken"`
}

// RegisterResponse représente la structure de la réponse à renvoyer
type RegisterResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
