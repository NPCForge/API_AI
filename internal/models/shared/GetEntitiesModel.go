package sharedModel

type Entity struct {
	Id       string `json:"id"`
	Checksum string `json:"checksum"`
}

type RequestGetEntities struct {
	Action string `json:"string"`
	Token  string `json:"token"`
}

type ResponseGetEntities struct {
	Entity []Entity `json:"entities"`
	Status string   `json:"status"`
}
