package models

type AuthResponse struct {
	Token string `json:"token"`
}

type ViewReponse struct {
	Page       int           `json:"page"`
	PerPage    int           `json:"perPage"`
	TotalItems int           `json:"totalItems"`
	TotalPages int           `json:"totalPages"`
	Items      []interface{} `json:"items"` // Utilise un type adapté selon ta structure d'élément
}

type BodyEntity struct {
	Token  string `json:"token"`
	Prompt string `json:"prompt"`
	Name   string `json:"name"`
}
