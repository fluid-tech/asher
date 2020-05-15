package generator

import "asher/internal/api/codebuilder/php/core"

type ModelGenerator struct {
	class *core.Class
}

/**
  Creates a new instance of this builder with a new core.Class
 */
func NewModelGenerator() *ModelGenerator {
	return &ModelGenerator{
		class: core.NewClass(),
	}
}