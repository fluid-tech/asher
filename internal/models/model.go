package models

type Model struct {
	Name        string     `json:"name"`
	Cols        []*Column  `json:"cols"`
	Relations   Relation   `json:"relations"`
	SoftDeletes bool       `json:"softDeletes"`
	Timestamps  bool       `json:"timestamps"`
	AuditCols   bool       `json:"auditCols"`
	Controller  Controller `json:"controller"`
}
