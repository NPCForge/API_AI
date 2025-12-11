package sharedModel

type EntityDiedRequest struct {
	Checksum string `json:"checksum"`
	Token    string `json:"token"`
}

type EntityDiedResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
