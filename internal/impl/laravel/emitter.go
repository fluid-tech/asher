package laravel

import (
	"asher/internal/api"
	"asher/internal/impl/laravel/5.8/handler/helper"
	"asher/internal/models"
)

type Emitter struct {
	api.Emitter
	activeHandlerRegister map[string]*api.Handler
	fileMap               map[string]api.EmitterFile
	version               string
}

func NewLaravelEmitter(version string) *Emitter {
	return &Emitter{version: version,
		fileMap:               map[string]api.EmitterFile{},
		activeHandlerRegister: map[string]*api.Handler{},
	}
}

func (e Emitter) Emit(value interface{}) {
	// todo make this work for everything apart from Model instances

	model := value.(models.Model)
	//arrayOfEmittedFiles, err := GetFromRegistry("cols").Handle(model.Name, model.Cols)
	// todo add if
	//GetFromRegistry("relations").Handle(model.Name, model.Cols)
	//GetFromRegistry("softDeletes").Handle(model.Name, model.SoftDeletes)
	//GetFromRegistry("timestamps").Handle(model.Name, model.Timestamps)
	GetFromRegistry("columns").Handle(model.Name, model.Cols)
	GetFromRegistry("controller").Handle(model.Name, model.Controller)
	GetFromRegistry("auditCols").Handle(model.Name, helper.NewAuditCol(model.AuditCols,
		model.SoftDeletes, model.Timestamps, pkColName(model.Cols)))

}

func (e Emitter) GetFileMap() map[string]api.EmitterFile {
	return e.fileMap
}

/*
 Finds the first instance of primary in cols and returns the name of the column. If none
 are found returns an empty string `""`.
 Parameters:
	- cols			[]*models.Column		The arr of cols to query
 Returns:
	- string		The name of the col
 Usage:
	pkColName(col)
*/
func pkColName(cols []*models.Column) string {
	for _, element := range cols {
		if element.Primary {
			return element.Name
		}
	}
	return ""
}
