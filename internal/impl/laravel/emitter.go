package laravel

import (
	"asher/internal/api"
	"asher/internal/models"
)

type Emitter struct {
	api.Emitter
	activeHandlerRegister map[string]*api.Handler
	fileMap               map[string]*api.EmitterFile
	version               string
}

func NewLaravelEmitter(version string) *Emitter {
	return &Emitter{version: version,
		fileMap:               map[string]*api.EmitterFile{},
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
	GetFromRegistry("auditCols").Handle(model.Name, model.AuditCols)

}

func (e Emitter) GetFileMap() map[string]*api.EmitterFile {
	return e.fileMap
}
