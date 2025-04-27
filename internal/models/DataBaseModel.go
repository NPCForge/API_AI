package models

type Entity struct {
	ID       int
	UserID   int
	Name     string
	Checksum string
	Prompt   string
	Created  string
}
