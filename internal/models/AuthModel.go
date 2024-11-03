package models

type AuthResponse struct {
	Token string `json:"token"`
}

type ViewReponse struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"perPage"`
	TotalItems int    `json:"totalItems"`
	TotalPages int    `json:"totalPages"`
	Items      []Item `json:"items"`
}

type Item struct {
	ID string `json:"id"`
}

type BodyEntity struct {
	Token  string `json:"token"`
	Prompt string `json:"prompt"`
	Name   string `json:"name"`
}

type DisconnectResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type RemoveResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
