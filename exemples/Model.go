package exemples

type Request struct {
	Action   string `json:"string"`
	Name     string `json:"name"`
	Prompt   string `json:"prompt"`
	Checksum string `json:"checksum"`
	Token    string `json:"token"`
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
