package handler

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/models"
	"errors"
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

	columnHandler.handleMigration(modelName, myColsArray)

	return []*api.EmitterFile{}, nil
}

func (columnHandler *ColumnHandler) handleMigration(identifier string, columnArr []models.Column) error {

	var statementsArr []core.SimpleStatement
	for _, singleColumn := range columnArr {
		statementsArr = append(
			statementsArr,
			columnHandler.handleSingleColumn(identifier, singleColumn),
		)
	}

	for _, stmt := range statementsArr {
		fmt.Println(stmt.String())
	}
	return nil

}

func (columnHandler *ColumnHandler) handleSingleColumn(modelName string, column models.Column) core.SimpleStatement {


	if column.Primary {
		//Handle PrimaryKey
		return columnHandler.handlePrimary(column)
	} else if column.ColType == "reference" {
		// Handle ForeignKey
		return columnHandler.handleForeign(column)
	} else {
		// Handle Other Columns
		return columnHandler.handleOther(column)
	}

}

func (columnHandler *ColumnHandler) handleFillable(modelName string, column models.Column) error {
	modelClass := context.GetFromRegistry("model").GetCtx(modelName).(*core.Class)
	if modelClass != nil {
		element, err := modelClass.FindInMembers("fillable")
		if err != nil {
			// todo return an err
			return err
		}
		arrayAssignment := (*element).(*core.ArrayAssignment)
		arrayAssignment.Rhs = append(arrayAssignment.Rhs, column.Name)
		return nil
	}
	return errors.New(fmt.Sprintf("model class %s not found", modelName))
}

func (columnHandler *ColumnHandler) handlePrimary(column models.Column) core.SimpleStatement {
	var generatedLine string
	if column.GenerationStrategy == "auto_increment" {
		generatedLine = "$table->"
		generatedLine += columnHandler.primaryKeyMethodNameGenerator(column.ColType)
		generatedLine += "('"
		generatedLine += column.Name
		generatedLine += "')"
	} else if column.GenerationStrategy == "uuid" {
		//$table->uuid('id')->primary();
		generatedLine = "$table->uuid('" + column.Name + "')->primary()"
	} else {
		panic("Input Error")
	}
	return core.SimpleStatement{
		SimpleStatement: generatedLine,
	}

}

func (columnHandler *ColumnHandler) handleOther(column models.Column) core.SimpleStatement {
	var generatedLine string
	generatedLine = "$table->" + columnHandler.ColTypeSwitcher(column.ColType, column.Name, column.Allowed)
	generatedLine += columnHandler.handleDefaultValue(column.DefaultVal) + columnHandler.handleNullable(column.Nullable) + columnHandler.handleUnique(column.Unique) + columnHandler.handleIndex(column.Index)
	return core.SimpleStatement{
		SimpleStatement: generatedLine,
	}
}

func (columnHandler *ColumnHandler) handleForeign(column models.Column) core.SimpleStatement {
	//$table->foreign('user_id')->references('id')->on('users')->onDelete('cascade');
	var sb string
	sb += "$table->foreign('"
	sb += column.Name
	sb += "')"
	//var tableName, columnName string
	var splitedArr []string
	splitedArr = strings.Split(column.Table, ":")
	sb += "->references('" + splitedArr[1] + "')"
	sb += "->on('" + splitedArr[0] + "')"
	sb += "->onDelete('" + column.OnDelete + "')"
	// TODO: check if the onDelete Value is sanitized of not
	// TODO: do more research on cascade ondelete
	return core.SimpleStatement{
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
	case "mediumInt":
		return "mediumIncrements"
	case "smallInteger":
		return "smallIncrements"
	case "tinyInt":
		return "tinyIncrements"
	case "bigInteger":
		return "bigIncrements"
	default:
		panic("Type not supported or invalid inputs")
	}
}

func (ColumnHandler *ColumnHandler) handleAllowedKeywordsToString(allowed []string) string {
	bldr := "'" + strings.Join(allowed, "', '") + "'"
	return "[" + bldr + "]"
}

/*
	This method will have all the keys defined by asher as valid input value and return
	its respective laravel method name
*/
func (columnHandler *ColumnHandler) ColTypeSwitcher(colType string, colName string, allowed []string) string {
	var columnNameBracket = "('" + colName + "')"
	switch colType {
	case "unsignedBigInteger":
		return "unsignedBigInteger" + columnNameBracket
	case "bigInteger":
		return "bigInteger" + columnNameBracket
	case "string":
		return "string" + columnNameBracket
	case "enum":
		return "enum" + "('" + colName + "', " + columnHandler.handleAllowedKeywordsToString(allowed) + ")"
	default:
		// TODO: Log this error and replace it with formatted error message.
		panic("not supported or wrong input in ColTypeSwitcher" + colType)
	}
}
