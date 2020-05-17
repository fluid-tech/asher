package generator

import (
	"asher/internal/api/codebuilder/php/core"
	"strings"
)

type QueryGenerator struct {

}

func NewQueryGenerator() *QueryGenerator {
	return &QueryGenerator{}
}

func (queryGenerator *QueryGenerator) Build() *core.Class {
	return &core.Class{}
}

func (queryGenerator *QueryGenerator) String() string {
	var builder strings.Builder
	return builder.String()
}

