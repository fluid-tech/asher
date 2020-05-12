package models

type Model struct{
	name        string
	cols        []Column
	relations   []Relation
	softDeletes bool
	timestamps  bool
	auditCols   bool
	controller  Controller
}

// todo add getters