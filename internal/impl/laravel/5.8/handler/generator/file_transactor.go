package generator

import (
	"asher/internal/api/codebuilder/php/core"
	"strings"
)

type FileTransactor struct {
	transactorGen *TransactorGenerator
}

func NewFileTransactor(transactorGen *TransactorGenerator) *FileTransactor {
	return &FileTransactor{transactorGen: transactorGen}
}

func (fileTransactor *FileTransactor) AddDefaults() *FileTransactor {
	fileTransactor.transactorGen.AppendImports([]string{BaseFileUploadHelperPath}).
		AddParentConstructorCallArgs(core.NewParameter(NewBaseFileUploadHelper)).
		AddTransactorMember(core.NewSimpleStatement(`private const BASE_PATH = "` +
			strings.ToLower(fileTransactor.transactorGen.identifier) + `"`)).
		AddTransactorMember(core.NewSimpleStatement(ImageValidationRules))
	return fileTransactor
}
