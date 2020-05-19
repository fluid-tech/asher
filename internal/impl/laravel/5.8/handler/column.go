package handler

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/impl/laravel/5.8/handler/helper"
	"asher/internal/models"
	"fmt"
	"strings"
)

type ColumnHandler struct {
	api.Handler
}

func NewColumnHandler() *ColumnHandler {
	return &ColumnHandler{}
}

func (columnHandler *ColumnHandler) Handle(modelName string, colsArr interface{}) ([]*api.EmitterFile, error) {

	myColsArray := colsArr.([]models.Column)

	tempMigration := api.EmitterFile(columnHandler.handleMigration(modelName, myColsArray))
	tempModel := api.EmitterFile(columnHandler.handleModel(modelName, myColsArray))

	return []*api.EmitterFile{&tempMigration, &tempModel}, nil
}

func (columnHandler *ColumnHandler) handleModel(modelName string, colArr []models.Column) *core.PhpEmitterFile {
	var modelGenerator = generator.NewModelGenerator().SetName(modelName)

	for _, singleColumn := range colArr {
		columnHandler.handleHidden(modelGenerator, singleColumn.Hidden, singleColumn.Name)
		columnHandler.handleGuarded(modelGenerator, singleColumn.Guarded, singleColumn.Name)
		columnHandler.handleValidation(modelGenerator, singleColumn.Validations, singleColumn.Name)
	}
	//fmt.Print(modelGenerator.Build().String())
	context.GetFromRegistry("model").AddToCtx(modelName, modelGenerator)
	modelGeneratorRef := api.Generator(modelGenerator)
	phpEmitter := core.NewPhpEmitterFile(modelName, api.ModelPath, &modelGeneratorRef, api.Model)
	return phpEmitter
}

func (columnHandler *ColumnHandler) handleMigration(identifier string, columnArr []models.Column) api.EmitterFile {
	var statementsArr []*core.SimpleStatement
	for _, singleColumn := range columnArr {
		statementsArr = append(
			statementsArr,
			columnHandler.handleSingleColumn(identifier, singleColumn),
		)

	}

	migrationGenerator := generator.NewMigrationGenerator().SetName(identifier).AddColumns(statementsArr)
	klass := migrationGenerator.Build()
	fmt.Println(klass)
	context.GetFromRegistry("migration").AddToCtx(identifier, migrationGenerator)
	modelGeneratorRef := api.Generator(migrationGenerator)
	phpEmitter := core.NewPhpEmitterFile(identifier, api.ModelPath, &modelGeneratorRef, api.Model)
	return phpEmitter
}

func (columnHandler *ColumnHandler) handleValidation(modelGenerator *generator.ModelGenerator, validations string, colName string) {
	if validations != "" {
		modelGenerator.AddCreateValidationRule(colName, validations)
		modelGenerator.AddUpdateValidationRule(colName, validations)
	}
}

func (columnHandler *ColumnHandler) handleSingleColumn(modelName string, column models.Column) *core.SimpleStatement {

	if column.Primary {
		//Handle PrimaryKey
		return columnHandler.handlePrimary(column.ColType, column.Name, column.GenerationStrategy)
	} else if column.ColType == "reference" {
		// Handle ForeignKey
		return columnHandler.handleForeign(column.Name, column.Table, column.OnDelete, column.Nullable)
	} else {
		// Handle Other Columns
		return columnHandler.handleOther(column)
	}

}

func (columnHandler *ColumnHandler) handleHidden(modelGenerator *generator.ModelGenerator, isHidden bool, colName string) {
	if isHidden {
		modelGenerator.AddHiddenField(colName)
	}
}

func (columnHandler *ColumnHandler) handleGuarded(modelGenerator *generator.ModelGenerator, isGuarded bool, colName string) {
	if isGuarded {
		modelGenerator.AddFillable(colName)
	}
}

func (columnHandler *ColumnHandler) handlePrimary(colType string, colName string, genStrat string) *core.SimpleStatement {
	var generatedLine string
	if genStrat == "auto_increment" {
		primaryKeyMethodName := columnHandler.primaryKeyMethodNameGenerator(colType)
		generatedLine = fmt.Sprintf("$table->%s('%s')", primaryKeyMethodName, colName)
	} else if genStrat == "uuid" {
		//$table->uuid('id')->primary();
		generatedLine = fmt.Sprintf("$table->uuid('%s')->primary()", colName)
	} else {
		panic("input Type does not match with the defined keywords (uuid, auto_increment)")
	}
	return &core.SimpleStatement{
		SimpleStatement: generatedLine,
	}

}

func (columnHandler *ColumnHandler) handleOther(column models.Column) *core.SimpleStatement {
	var generatedLine string
	colTypeVal := columnHandler.ColTypeSwitcher(column.ColType, column.Name, column.Allowed)
	defaultVal := columnHandler.handleDefaultValue(column.DefaultVal)
	nullableVal := columnHandler.handleNullable(column.Nullable)
	uniqueVal := columnHandler.handleUnique(column.Unique)
	indexVal := columnHandler.handleIndex(column.Index)
	generatedLine = fmt.Sprintf("$table->%s%s%s%s%s", colTypeVal, defaultVal, nullableVal, uniqueVal, indexVal)
	return &core.SimpleStatement{
		SimpleStatement: generatedLine,
	}
}

func (columnHandler *ColumnHandler) handleForeign(colName string, colTable string, colOnDelete string, isNullable bool) *core.SimpleStatement {
	//$table->foreign('user_id')->references('id')->on('users')->onDelete('cascade');
	var sb string
	var splitedArr []string
	splitedArr = strings.Split(colTable, ":")
	sb = fmt.Sprintf("$table->foreign('%s')->references('%s')->on('%s')->onDelete('%s')%s", colName, splitedArr[1], splitedArr[0], colOnDelete, columnHandler.handleNullable(isNullable))

	// TODO: check if the onDelete Value is sanitized of not
	// TODO: do more research on cascade ondelete will it be mandatory or not
	return &core.SimpleStatement{
		SimpleStatement: sb,
	}
}

func (columnHandler *ColumnHandler) handleUnique(isUnique bool) string {
	if isUnique {
		return "->unique()"
	}
	return ""
}

func (columnHandler *ColumnHandler) handleDefaultValue(defaultVal string) string {
	if defaultVal != "" {
		return "->default('" + defaultVal + "')"
	}
	return ""
}

func (columnHandler *ColumnHandler) handleIndex(isIndex bool) string {
	if isIndex {
		return "->index()"
	}
	return ""
}

func (columnHandler *ColumnHandler) handleNullable(isNullable bool) string {
	if isNullable {
		return "->nullable()"
	}
	return ""
}

func (columnHandler *ColumnHandler) primaryKeyMethodNameGenerator(colType string) string {
	switch colType {
	case "integer":
		return "increments"
	case "mediumInteger":
		return "mediumIncrements"
	case "smallInteger":
		return "smallIncrements"
	case "tinyInteger":
		return "tinyIncrements"
	case "bigInteger":
		return "bigIncrements"
	default:
		panic("Type not supported or invalid inputs")
	}
}

/*
	This method will have all the keys defined by asher as valid input value and return
	its respective laravel method name
*/
func (columnHandler *ColumnHandler) ColTypeSwitcher(colType string, colName string, allowed []string) string {
	switch colType {
	// TODO : Add more column types here
	case "unsignedBigInteger":
		return helper.UnsignedBigInteger(colName)
	case "bigInteger":
		return helper.BigInteger(colName)
	case "unsignedInteger":
		return helper.UnsignedInteger(colName)
	case "integer":
		return helper.Integer(colName)
	case "unsignedTinyInteger":
		return helper.UnsignedTinyInteger(colName)
	case "tinyInteger":
		return helper.TinyInteger(colName)
	case "unsignedMediumInteger":
		return helper.UnsignedMediumInteger(colName)
	case "mediumInteger":
		return helper.MediumInteger(colName)
	case "string":
		return helper.String(colName)
	case "boolean":
		return helper.Boolean(colName)
	case "char":
		return helper.Char(colName)
	case "date":
		return helper.Date(colName)
	case "double":
		return helper.Double(colName)
	case "float":
		return helper.Float(colName)
	case "enum":
		return helper.Enum(colName, allowed)
	case "set":
		return helper.Set(colName, allowed)
	default:
		// TODO: Log this error and replace it with formatted error message.
		panic("not supported or wrong input in ColTypeSwitcher :- " + colType)
	}
}
