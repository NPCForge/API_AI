package httpModels

// RegisterRequest représente la structure des données attendues dans la requête
type RegisterRequest struct {
	Token  string `json:"BearerToken"`
	Name   string `json:"name"`
	Prompt string `json:"prompt"`
}

// RegisterResponse représente la structure de la réponse à renvoyer
type RegisterResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Private string `json:"private"`
}
