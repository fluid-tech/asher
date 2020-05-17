package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
)

type ModelGenerator struct {
	api.Generator
	classBuilder          interfaces.Class
	fillables             []string
	hidden                []string
	timestamps            bool
	createValidationRules map[string]string
	updateValidationRules map[string]string
}

/**
Creates a new instance of this generator with a new interfaces.Class
Returns:
	- instance of generator object
*/
func NewModelGenerator() *ModelGenerator {
	return &ModelGenerator{
		classBuilder:          builder.NewClassBuilder(),
		fillables:             []string{},
		hidden:                []string{},
		timestamps:            false,
		createValidationRules: map[string]string{},
		updateValidationRules: map[string]string{},
	}
}

/**
Adds the validation rule in createValidationRules array which will is in form of associative array.
Parameters:
	- colName: this will be the Name of Column
	- colRule: This will be rule/constraint imposed on the specified passed column.
Returns:
	- instance of the generator object
Example:
	- AddCreateValidationRule('student_name', 'max:255|string')
*/
func (modelGenerator *ModelGenerator) AddCreateValidationRule(colName string, colRule string) *ModelGenerator {
	modelGenerator.createValidationRules = make(map[string]string)
	modelGenerator.createValidationRules[colName] = colRule
	return modelGenerator
}

/**
Adds the validation rule in updateValidationRules array which will is in form of associative array.
Parameters:
	- colName: this will be the Name of Column
	- colRule: This will be rule/constraint imposed on the specified passed column.
Returns:
	- instance of the generator object
Example:
	- AddUpdateValidationRule('student_name', 'string|max:255')
*/
func (modelGenerator *ModelGenerator) AddUpdateValidationRule(colName string, colRule string) *ModelGenerator {
	modelGenerator.updateValidationRules = make(map[string]string)
	modelGenerator.updateValidationRules[colName] = colRule
	return modelGenerator
}

/**
 Set's the name of the model class.
 Parameters:
	- tableName: the name of table specified in asher config.
 Returns:
	- instance of the generator object
 Example:
	- SetName('student_allotments')
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
 Example:
	- AddFillable('student_name')
*/
func (modelGenerator *ModelGenerator) AddFillable(columnName string) *ModelGenerator {
	modelGenerator.fillables = append(modelGenerator.fillables, `"`+columnName+`"`)
	return modelGenerator
}

/**
 Adds the given column to hidden fields of this model
 Parameters:
	- columnName: the name of the column to add
 Returns:
	- instance of the generator object
 Example:
	- AddHiddenField('student_name')
*/
func (modelGenerator *ModelGenerator) AddHiddenField(columnName string) *ModelGenerator {
	modelGenerator.hidden = append(modelGenerator.hidden, `"`+columnName+`"`)
	return modelGenerator
}

/**
 Control whether to set timestamps in the model of not
 Returns:
	- instance of the generator object
 Example:
	- SetTimestamps(true)
*/
func (modelGenerator *ModelGenerator) SetTimestamps(flag bool) *ModelGenerator {
	modelGenerator.timestamps = flag
	return modelGenerator
}

/**
 Builds the corresponding model from the given ingredients of input.
 Note: It returns a new core.Class object every time it's called.
 Returns:
 	- The corresponding model core.Class from the given ingredients of input.
 */
func (modelGenerator *ModelGenerator) Build() *core.Class {
	modelGenerator.classBuilder = modelGenerator.classBuilder.SetPackage("App").AddImports([]string{
		`Illuminate\Database\Eloquent\Model`,
	}).SetExtends("Model")

	if len(modelGenerator.fillables) > 0 {
		fillableArray := api.TabbedUnit(core.NewArrayAssignment("protected", "fillable",
			modelGenerator.fillables))
		modelGenerator.classBuilder = modelGenerator.classBuilder.AddMember(&fillableArray)
	}

	if len(modelGenerator.hidden) > 0 {
		hiddenArray := api.TabbedUnit(core.NewArrayAssignment("protected", "visible",
			modelGenerator.hidden))
		modelGenerator.classBuilder = modelGenerator.classBuilder.AddMember(&hiddenArray)
	}

	if modelGenerator.timestamps {
		timestamps := api.TabbedUnit(core.NewVarAssignment("public", "timestamps", "true"))
		modelGenerator.classBuilder = modelGenerator.classBuilder.AddMember(&timestamps)
	}

	if len(modelGenerator.createValidationRules) > 0 {
		createFunction := getValidationRulesFunction("createValidationRules",
			modelGenerator.createValidationRules)
		modelGenerator.classBuilder = modelGenerator.classBuilder.AddFunction(createFunction)
	}

	if len(modelGenerator.updateValidationRules) > 0 {
		updateFunction := getValidationRulesFunction("updateValidationRules",
			modelGenerator.updateValidationRules)
		modelGenerator.classBuilder = modelGenerator.classBuilder.AddFunction(updateFunction)
	}

	return modelGenerator.classBuilder.GetClass()
}

/**
 Implementation of the base Generator to return string of this model.
 Returns:
	- string representation of this mode.
 */
func (modelGenerator *ModelGenerator) String() string {
	return modelGenerator.Build().String()
}


/**
 A helper function to generate a ReturnArray for rules with the given method name.
 Parameters:
	- functionName: name of the function that returns these validation rules.
	- rules: a map of columns and their validation rules.
 Returns:
	- instance of core.Function for the given input.
 Example:
	- getValidationRulesFunction('createValidationRules', map[string]string{'name':'max:255'})
*/
func getValidationRulesFunction(functionName string, rules map[string]string) *core.Function {
	returnArray := api.TabbedUnit(core.NewReturnArrayFromMap(rules))
	function := builder.NewFunctionBuilder().SetName(functionName).
		SetVisibility("public").AddStatement(&returnArray).GetFunction()
	return function
}