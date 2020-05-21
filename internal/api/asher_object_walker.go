package api

import (
	"asher/internal/models"
)

type AsherObjectWalker struct {
	emitter     Emitter      // the Emitter this instance of Walker uses
	asherObject models.Asher // the object this instance iterates over
	writer      *Writer
}

/**
Constructs a new AsherObjectWalker with the given args
*/
func NewAsherObjectWalker(asherObject models.Asher, emitter Emitter, registry WriterHandlerRegistry) *AsherObjectWalker {
	return &AsherObjectWalker{
		emitter:     emitter,
		asherObject: asherObject,
		writer:      NewWriter(registry),
	}
}

/**
Walks over the asher object provided and triggers callbacks
*/
func (walker AsherObjectWalker) Walk() {
	walker.walkModels()
	emitterFiles := walker.emitter.GetFiles()
	// TODO: add topological sort to emitterFiles to avoid problems with dependencies
	walker.writer.Walk(emitterFiles)
}

func (walker AsherObjectWalker) walkModels() {
	for _, model := range walker.asherObject.Models {
		walker.emitter.Emit(model)
	}
}
