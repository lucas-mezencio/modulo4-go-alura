package models

type Personality struct {
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalities []Personality
