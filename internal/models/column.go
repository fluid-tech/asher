package models

type Column struct {
	Name               string   `json:"name"`               // name of the col
	ColType            string   `json:"colType"`            // could be integer, float, reference, enum etc
	GenerationStrategy string   `json:"generationStrategy"` // to be used when the col is a primary, includes auto_increment, uuid etc
	DefaultVal         string   `json:"defaultVal"`         // the default val to be assigned
	Table              string   `json:"table"`              // the table this maps a foreign key from
	Validations        string   `json:"validations"`        // a set of validation rules separated by |
	Index              bool     `json:"index"`				// should this col be indexed
	Allowed            []string `json:"allowed"`   			// a set of allowed values to be used in cases of enums
	Invisible          bool     `json:"invisible"` 			// indicates whether this col is presented when jsonified
}
