package handler

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/impl/laravel/5.8/handler/helper"
	"asher/internal/models"
	"errors"
	"fmt"
	"strings"
)

type ColumnHandler struct {
	api.Handler
}

const FkFmt = "$table->foreign('%s')->references('%s')->on('%s')->onDelete('%s')%s"
const UuidFmt = "$table->uuid('%s')->primary()"

func NewColumnHandler() *ColumnHandler {
	return &ColumnHandler{}
}

func (columnHandler *ColumnHandler) Handle(modelName string, colsArr interface{}) ([]api.EmitterFile, error) {

	myColsArray := colsArr.([]models.Column)

	tempMigration, err := columnHandler.handleMigration(modelName, myColsArray)
	if err != nil {
		return nil, err
	}
	tempModel, err := columnHandler.handleModel(modelName, myColsArray)
	if err != nil {
		return nil, err
	}
	return []api.EmitterFile{tempMigration, tempModel}, nil
}

func (columnHandler *ColumnHandler) handleModel(modelName string, colArr []models.Column) (*core.PhpEmitterFile, error) {
	var modelGenerator = generator.NewModelGenerator().SetName(modelName)

	for _, singleColumn := range colArr {
		columnHandler.handleHidden(modelGenerator, singleColumn.Hidden, singleColumn.Name)
		columnHandler.handleGuarded(modelGenerator, singleColumn.Guarded, singleColumn.Name)
		columnHandler.handleValidation(modelGenerator, singleColumn.Validations, singleColumn.Name, modelName)
	}
	context.GetFromRegistry(context.Model).AddToCtx(modelName, modelGenerator)

	return core.NewPhpEmitterFile(modelName, api.ModelPath, modelGenerator, api.Model), nil
}

func (columnHandler *ColumnHandler) handleMigration(identifier string, columnArr []models.Column) (*core.PhpEmitterFile,
	error) {
	var statementsArr []*core.SimpleStatement
	for _, singleColumn := range columnArr {
		stmt, err := columnHandler.handleSingleColumn(identifier, singleColumn)
		if err != nil {
			return nil, err
		}
		statementsArr = append(
			statementsArr,
			stmt,
		)
	}

	migrationGenerator := generator.NewMigrationGenerator().SetName(identifier).AddColumns(statementsArr)
	context.GetFromRegistry(context.Migration).AddToCtx(identifier, migrationGenerator)

	return core.NewPhpEmitterFile(identifier, api.MigrationPath, migrationGenerator, api.Model), nil
}

func (columnHandler *ColumnHandler) handleValidation(modelGenerator *generator.ModelGenerator, validations string,
	colName string, modelName string) {
	if validations != "" {
		modelGenerator.AddCreateValidationRule(colName, validations, modelName)
		modelGenerator.AddUpdateValidationRule(colName, validations, modelName)
	}
}

func (columnHandler *ColumnHandler) handleSingleColumn(modelName string, column models.Column) (*core.SimpleStatement,
	error) {

	if column.Primary {
		//Handle PrimaryKey
		stmt, err := columnHandler.handlePrimary(column.ColType, column.Name, column.GenerationStrategy)
		if err != nil {
			return nil, err
		}
		return stmt, nil
	} else if column.ColType == "reference" {
		// Handle ForeignKey
		return columnHandler.handleForeign(column.Name, column.Table, column.OnDelete, column.Nullable), nil
	} else {
		// Handle Other Columns
		return columnHandler.handleOther(column)
	}

}

func (columnHandler *ColumnHandler) handleHidden(modelGenerator *generator.ModelGenerator, isHidden bool,
	colName string) {
	if isHidden {
		modelGenerator.AddHiddenField(colName)
	}
}

func (columnHandler *ColumnHandler) handleGuarded(modelGenerator *generator.ModelGenerator, isGuarded bool,
	colName string) {
	if isGuarded {
		modelGenerator.AddFillable(colName)
	}
}

/*
If genStrat == auto_increment Then "$table->increments('colName')"
if genStrat == uuid Then "$table->uuid('id')->primary()"
else Return an Error
*/
func (columnHandler *ColumnHandler) handlePrimary(colType string, colName string,
	genStrat string) (*core.SimpleStatement, error) {
	var generatedLine string
	if genStrat == "auto_increment" {
		primaryKeyMethodName, err := helper.PrimaryKeyMethodNameGenerator(colType)
		if err != nil {
			return nil, err
		}
		generatedLine = fmt.Sprintf("$table->%s('%s')", primaryKeyMethodName, colName)
	} else if genStrat == "uuid" {
		//$table->uuid('id')->primary();
		generatedLine = fmt.Sprintf(UuidFmt, colName)
	} else {
		return nil, errors.New("input Type does not match with the defined keywords (uuid, auto_increment)")
	}
	return &core.SimpleStatement{
		SimpleStatement: generatedLine,
	}, nil

}

func (columnHandler *ColumnHandler) handleOther(column models.Column) (*core.SimpleStatement, error) {
	var generatedLine string
	colTypeVal, err := helper.ColTypeSwitcher(column.ColType, column.Name, column.Allowed)
	if err != nil {
		return nil, err
	}
	defaultVal := columnHandler.handleDefaultValue(column.DefaultVal)
	nullableVal := columnHandler.handleNullable(column.Nullable)
	uniqueVal := columnHandler.handleUnique(column.Unique)
	indexVal := columnHandler.handleIndex(column.Index)
	generatedLine = fmt.Sprintf("$table->%s%s%s%s%s", colTypeVal, defaultVal, nullableVal, uniqueVal, indexVal)
	return &core.SimpleStatement{
		SimpleStatement: generatedLine,
	}, nil
}

func (columnHandler *ColumnHandler) handleForeign(colName string, colTable string, colOnDelete string,
	isNullable bool) *core.SimpleStatement {
	//$table->foreign('user_id')->references('id')->on('users')->onDelete('cascade');
	var sb string
	var splitedArr []string
	splitedArr = strings.Split(colTable, ":")
	sb = fmt.Sprintf(FkFmt, colName, splitedArr[1], splitedArr[0], colOnDelete, columnHandler.handleNullable(isNullable))

	// TODO: check if the onDelete Value is sanitized of not
	// TODO: do more research on cascade ondelete will it be mandatory or not
	return &core.SimpleStatement{
		SimpleStatement: sb,
	}
}

/**
 An UniqueHandler function to generate a String for column if the parameter is set TRUE.
 Parameters:
	- isIndex: if this value is true then the column is to be unique
 Returns:
	- "->unique()" string if the parameter is TRUE else BLANK ""
 Example:
	- handleUnique(true)
*/
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

/**
 An IndexHandler function to generate a String for column if the parameter is set TRUE.
 Parameters:
	- isIndex: if this value is true then the column is to be indexed
 Returns:
	- "->index()" string if the parameter is TRUE else BLANK ""
 Example:
	- handleIndex(true)
*/
func (columnHandler *ColumnHandler) handleIndex(isIndex bool) string {
	if isIndex {
		return "->index()"
	}
	return ""
}

/**
 A NullableHandler function to generate a String for column if the parameter is set TRUE.
 Parameters:
	- isNullable: if this value is true then the column is to be Nullable
 Returns:
	- "->nullable()" string if the parameter is TRUE else BLANK ""
 Example:
	- handleNullable(true)
*/
func (columnHandler *ColumnHandler) handleNullable(isNullable bool) string {
	if isNullable {
		return "->nullable()"
	}
	return ""
}
