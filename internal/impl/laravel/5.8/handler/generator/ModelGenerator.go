package generator

import (
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"github.com/iancoleman/strcase"
)

type ModelGenerator struct {
	class interfaces.Class
	fillables []string
	hidden []string
}

/**
 Creates a new instance of this builder with a new core.Class
 */
func NewModelGenerator() *ModelGenerator {
	return &ModelGenerator{
		class: builder.NewClassBuilder(),
	}
}

/**
 Set's the name of the model class.
 Parameters:
	- tableName: the name of table specified in asher config.
 Returns:
	- instance of the generator object
 */
func (modelGenerator *ModelGenerator) SetName(tableName string) *ModelGenerator {
	className := strcase.ToCamel(tableName)
	modelGenerator.SetName(className)
	return modelGenerator
}

/**
 Adds the given column to fillable fields of this model
 Parameters:
	- columnName: the name of the column to add
 Returns:
	- instance of the generator object
 */
func (modelGenerator *ModelGenerator) AddFillable(columnName string) *ModelGenerator {
	modelGenerator.fillables = append(modelGenerator.fillables, columnName)
	return modelGenerator
}

/**
 Adds the given column to hidden fields of this model
 Parameters:
	- columnName: the name of the column to add
 Returns:
	- instance of the generator object
*/
func (modelGenerator *ModelGenerator) AddHiddenField(columnName string) *ModelGenerator {
	modelGenerator.hidden = append(modelGenerator.hidden, columnName)
	return modelGenerator
}
