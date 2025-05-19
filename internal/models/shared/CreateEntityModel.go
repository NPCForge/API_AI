package sharedModel

type RequestCreateEntity struct {
	Action   string `json:"string"`
	Name     string `json:"name"`
	Prompt   string `json:"prompt"`
	Checksum string `json:"checksum"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type ResponseCreateEntity struct {
	Id       string `json:"id"`
	Checksum string `json:"checksum"`
	Message  string `json:"message"`
	Status   string `json:"status"`
}
