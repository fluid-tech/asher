package writer

import (
	"asher/internal/api"
)

type MigrationWriterHandler struct {
	api.WriterHandler
	migrationFileName	string
}

func NewMigrationWriterHandler() *MigrationWriterHandler {
	return &MigrationWriterHandler{
		migrationFileName: "",
	}
}

func (writerHandler *MigrationWriterHandler) Handle(emitterFile api.EmitterFile) int {

}

func (writerHandler *MigrationWriterHandler) prepare(emitterFile api.EmitterFile) {

}