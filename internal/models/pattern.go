package models

type Pattern struct {
	PatternType string   `json:"patternType"`
	Protected   bool     `json:"bool"` // wrap with auth middleware
	NestedWith  []string `json:"nestedWith"`
	Type        string   `json:"type"` // default, file etc
}
