package api

import "asher/internal/api/codebuilder/php/core"

type EmitterFile struct {
	FileName string      // name of the file
	Path     string      // path to store it in
	klass    *core.Class // class that must be stringified
	FileType int         // 0 - migration, 1 - model, 2- mutator, 3-transactor
}

func NewEmitterFile(fileName string, path string, class *core.Class, fileType int) *EmitterFile {
	return &EmitterFile{
		FileName: fileName,
		Path:     path,
		klass:    class,
		FileType: fileType,
	}
}
