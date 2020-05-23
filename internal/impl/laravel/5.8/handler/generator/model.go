package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
	"strings"
)

type ModelGenerator struct {
	api.Generator
	classBuilder          interfaces.Class
	fillables             []string
	hidden                []string
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
func (modelGenerator *ModelGenerator) AddCreateValidationRule(colName string, colRule string, modelName string) *ModelGenerator {

	returnString := "[ "
	var ruleArray = strings.Split(colRule, "|")
	for i := 0; i < len(ruleArray); i++ {
		if strings.HasPrefix(ruleArray[i], "unique:") {
			tableDataSplitter := strings.Split(ruleArray[i], ",")
			//tableName := strings.TrimPrefix(tableDataSplitter[0], "unique:")
			if len(tableDataSplitter) == 1 {
				ruleArray[i] = `'` + ruleArray[i] + `,` + colName + `'`
			} else {
				ruleArray[i] = `'` + ruleArray[i] + `'`
			}
		} else if ruleArray[i] == "unique" {
			ruleArray[i] = `'` + ruleArray[i] + ":" + modelName + "," + colName + `'`
		} else {
			ruleArray[i] = `'` + ruleArray[i] + `'`
		}
	}

	returnString = returnString + strings.Join(ruleArray, ", ")

	modelGenerator.createValidationRules[colName] = returnString + ` ]`
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
func (modelGenerator *ModelGenerator) AddUpdateValidationRule(colName string, colRule string, modelName string) *ModelGenerator {

	returnString := "[ "
	var ruleArray = strings.Split(colRule, "|")
	for i := 0; i < len(ruleArray); i++ {
		if strings.HasPrefix(ruleArray[i], "unique:") {
			tableDataSplitter := strings.Split(ruleArray[i], ",")
			tableName := strings.TrimPrefix(tableDataSplitter[0], "unique:")
			if len(tableDataSplitter) == 1 {
				ruleArray[i] = `'` + ruleArray[i] + `,` + colName + `,' . $row_ids['` + tableName + `']`
			} else {
				ruleArray[i] = `'` + ruleArray[i] + `,' . $row_ids['` + tableName + `']`
			}
		} else if ruleArray[i] == "unique" {
			ruleArray[i] = `'` + ruleArray[i] + ":" + modelName + "," + colName + `,' . $row_ids['` + modelName + `']`
		} else {
			ruleArray[i] = `'` + ruleArray[i] + `'`
		}
	}

	returnString = returnString + strings.Join(ruleArray, ", ")

	modelGenerator.updateValidationRules[colName] = returnString + ` ]`
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
Builds the corresponding model from the given ingredients of input.
Note: It returns a new core.Class object every time it's called.
Returns:
	- The corresponding model core.Class from the given ingredients of input.
*/
func (modelGenerator *ModelGenerator) Build() *core.Class {
	modelGenerator.classBuilder = modelGenerator.classBuilder.SetPackage("App").AddImport(
		`Illuminate\Database\Eloquent\Model`,
	).SetExtends("Model")

	if len(modelGenerator.fillables) > 0 {
		fillableArray := api.TabbedUnit(core.NewArrayAssignment("protected", "fillable",
			modelGenerator.fillables))
		modelGenerator.classBuilder.AddMember(fillableArray)
	}

	if len(modelGenerator.hidden) > 0 {
		hiddenArray := api.TabbedUnit(core.NewArrayAssignment("protected", "visible",
			modelGenerator.hidden))
		modelGenerator.classBuilder.AddMember(hiddenArray)
	}

	if len(modelGenerator.createValidationRules) > 0 {
		createFunction := getCreateValidationRulesFunction(modelGenerator.createValidationRules)
		modelGenerator.classBuilder.AddFunction(createFunction)
	}

	if len(modelGenerator.updateValidationRules) > 0 {
		updateFunction := getUpdateValidationRulesFunction(modelGenerator.updateValidationRules)
		modelGenerator.classBuilder.AddFunction(updateFunction)
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
	- rules: a map of columns and their validation rules.
 Returns:
	- instance of core.Function for the given input.
 Example:
	- getValidationRulesFunction(map[string]string{'name':'max:255'})
*/
func getUpdateValidationRulesFunction(rules map[string]string) *core.Function {
	returnArray := api.TabbedUnit(core.NewReturnArrayFromMapRaw(rules))
	function := builder.NewFunctionBuilder().SetName("updateValidationRules").
		SetVisibility("public").AddStatement(returnArray).SetStatic(true).AddArgument("$row_ids").GetFunction()
	return function
}

/**
 A helper function to generate a ReturnArray for rules with the given method name.
 Parameters:
	- rules: a map of columns and their validation rules.
 Returns:
	- instance of core.Function for the given input.
 Example:
	- getValidationRulesFunction(map[string]string{'name':'max:255'})
*/
func getCreateValidationRulesFunction(rules map[string]string) *core.Function {
	returnArray := api.TabbedUnit(core.NewReturnArrayFromMapRaw(rules))
	function := builder.NewFunctionBuilder().SetName("createValidationRules").
		SetVisibility("public").AddStatement(returnArray).SetStatic(true).GetFunction()
	return function
}
