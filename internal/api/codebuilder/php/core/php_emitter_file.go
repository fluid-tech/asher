package core

import (
	"asher/internal/api"
	"fmt"
	"strings"
)

type PhpEmitterFile struct {
	api.EmitterFile
	fileName string
	path     string
	fileType int
	content  *api.Generator
}

func NewPhpEmitterFile(name string, path string, units *api.Generator, fileType int) *PhpEmitterFile {
	return &PhpEmitterFile{
		fileName: name,
		path:     path,
		fileType: fileType,
		content:  units,
	}
}

func (f *PhpEmitterFile) FileName() string {
	return f.fileName
}

func (f *PhpEmitterFile) Path() string {
	return f.path
}

func (f *PhpEmitterFile) Content() *api.Generator {
	return f.content
}

func (f *PhpEmitterFile) FileType() int {
	return f.fileType
}

func (p *PhpEmitterFile) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "<?php\n")
	fmt.Fprint(&builder, (*p.Content()).String(), "\n")
	return builder.String()
}

