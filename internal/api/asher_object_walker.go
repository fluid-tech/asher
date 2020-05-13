package api

import (
	"asher/internal/models"
	"fmt"
)

type AsherObjectWalker struct {
	emitter     Emitter      // the Emitter this instance of Walker uses
	asherObject models.Asher // the object this instance iterates over
}

func NewAsherObjectWalker(asherObject models.Asher, emitter Emitter) *AsherObjectWalker {
	return &AsherObjectWalker{emitter: emitter, asherObject: asherObject}
}

func (walker AsherObjectWalker) Register(emitter Emitter) {
	walker.emitter = emitter
}

func (walker AsherObjectWalker) Walk() {
	walker.walkModels()
	fileData := walker.emitter.GetFileMap()
	// todo write to file
	fmt.Print(fileData)
}

func (walker AsherObjectWalker) walkModels() {
	for _, model := range walker.asherObject.Models {
		walker.emitter.Emit(model)
	}
}
