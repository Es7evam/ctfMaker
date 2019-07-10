package models

// Challenge struct stores a challenge options.
// It has Name, Description (Desc), Value, Flag and Category (Type)
type Challenge struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Value int    `json:"value"`
	Flag  string `json:"flag"`
	Type  string `json:"type"`
}
