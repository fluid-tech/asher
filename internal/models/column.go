package models

type Column struct{
	name string // name of the col
	colType string // could be integer, float, reference, enum etc
	generationStrategy string // to be used when the col is a primary, includes auto_increment, uuid etc
	defaultVal string // the default val to be assigned
	table string // the table this maps a foreign key from
	validations string // a set of validation rules separated by |
	index bool
	allowed []string // a set of allowed values to be used in cases of enums
	invisible bool // indicates whether this col is presented when jsonified
}
// todo add getters