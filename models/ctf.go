package models

// CTF model describes the challenge structure
type CTF struct {
	Name   string      `json:"name"`
	Challs []Challenge `json:"challs"`
	Tags   []string    `json:"tags"`
}
