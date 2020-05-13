package models

type Controller struct {
	Rest    bool    `json:"rest"`
	Mvc     bool    `json:"mvc"`
	Pattern Pattern `json:"pattern"`
}
