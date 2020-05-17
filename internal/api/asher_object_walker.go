package api

import (
	"asher/internal/models"
)

type AsherObjectWalker struct {
	emitter     Emitter      // the Emitter this instance of Walker uses
	asherObject models.Asher // the object this instance iterates over
}

/**
Constructs a new AsherObjectWalker with the given args
*/
func NewAsherObjectWalker(asherObject models.Asher, emitter Emitter) *AsherObjectWalker {
	return &AsherObjectWalker{emitter: emitter, asherObject: asherObject}
}

/**
Walks over the asher object provided and triggers callbacks
*/
func (walker AsherObjectWalker) Walk() {
	walker.walkModels()
	_ = walker.emitter.GetFileMap()
	// todo write to file

}

func (walker AsherObjectWalker) walkModels() {
	for _, model := range walker.asherObject.Models {
		walker.emitter.Emit(model)
	}
}
