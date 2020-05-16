package generator

import (
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
)

type ModelGenerator struct {
	classBuilder interfaces.Class
	fillables    []string
	hidden       []string
	timestamps   bool
}

// TODO : Add a method to edit/append to an existing core.Class

/**
 Creates a new instance of this builder with a new core.Class
 */
func NewModelGenerator() *ModelGenerator {
	return &ModelGenerator{
		classBuilder: builder.NewClassBuilder(),
	}
}

/**
 Set's the name of the model classBuilder.
 Parameters:
	- tableName: the name of table specified in asher config.
 Returns:
	- instance of the generator object
 */
func (modelGenerator *ModelGenerator) SetName(tableName string) *ModelGenerator {
	className := strcase.ToCamel(tableName)
	modelGenerator.classBuilder.SetName(className)
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
	modelGenerator.fillables = append(modelGenerator.fillables, `"` + columnName + `"`)
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
	modelGenerator.hidden = append(modelGenerator.hidden, `"` + columnName + `"`)
	return modelGenerator
}

/**
 Control whether to set timestamps in the model of not
 Returns:
	- instance of the generator object
*/
func (modelGenerator *ModelGenerator) SetTimestamps(flag bool) *ModelGenerator {
	modelGenerator.timestamps = flag
	return modelGenerator
}

func (modelGenerator *ModelGenerator) Build() *core.Class {
	modelGenerator.classBuilder = modelGenerator.classBuilder.SetPackage("App").AddImports([]string{
		`Illuminate\Database\Eloquent\Model`,
	})

	if len(modelGenerator.fillables) > 0 {
		fillableArray := core.TabbedUnit(core.NewArrayAssignment("protected", "fillable",
			modelGenerator.fillables))
		modelGenerator.classBuilder = modelGenerator.classBuilder.AddMember(&fillableArray)
	}

	if len(modelGenerator.hidden) > 0 {
		hiddenArray := core.TabbedUnit(core.NewArrayAssignment("protected", "visible",
			modelGenerator.hidden))
		modelGenerator.classBuilder = modelGenerator.classBuilder.AddMember(&hiddenArray)
	}

	if modelGenerator.timestamps {
		timestamps := core.TabbedUnit(core.NewVarAssignment("public","timestamps", "true"))
		modelGenerator.classBuilder = modelGenerator.classBuilder.AddMember(&timestamps)
	}

	return modelGenerator.classBuilder.GetClass()
}

