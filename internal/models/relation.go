package models

type Relation struct {
	HasMany []string `json:"hasMany"`
	HasOne []string  `json:"hasOne"`
}