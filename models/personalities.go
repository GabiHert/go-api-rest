package models

type Personality struct {
	Id    int    `json:"id"`
	Name_ string `json:"name_"`
	Story string `json:"story"`
}

var Personalities []Personality
