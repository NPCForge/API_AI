package sharedModel

type RemoveEntityRequest struct {
	Action   string `json:"string"`
	Checksum string `json:"checksum"`
	Token    string `json:"token"`
}

type RemoveEntityResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
