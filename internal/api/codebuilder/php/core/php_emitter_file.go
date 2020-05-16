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
	content  []*api.TabbedUnit
}

func NewPhpEmitterFile(name string, path string, units []*api.TabbedUnit, fileType int) *PhpEmitterFile {
	return &PhpEmitterFile{
		fileName:    name,
		path:        path,
		fileType:    fileType,
		content:     units,
	}
}

func (f *PhpEmitterFile) FileName() string {
	return f.fileName
}

func (f *PhpEmitterFile) Path() string {
	return f.path
}

func (f *PhpEmitterFile) Content() []*api.TabbedUnit {
	return f.content
}

func (f *PhpEmitterFile) FileType() int {
	return f.fileType
}

func (p *PhpEmitterFile) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "<?php\n")
	for _, element := range p.content{
		fmt.Fprint(&builder, (*element).String(), "\n")
	}
	return builder.String()
}
