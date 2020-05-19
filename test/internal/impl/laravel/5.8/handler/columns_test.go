package handler

import (
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/models"
	"fmt"
	"reflect"
	"testing"
)

type IN struct {
	tableName        string
	columnInputArray []models.Column
}

type OUT struct {
	expectedOutputFillable []string
	expectedOutputHidden   []string
}

func getInput(tableName string, columnInputArray []models.Column) IN {
	return IN{tableName, columnInputArray}
}

func expectedOutput(fillableExpected []string, hiddenExpected []string) OUT {
	return OUT{fillableExpected, hiddenExpected}
}

func Test_Columns(t *testing.T) {

	var columnTestObject = []*struct {
		in  IN
		out OUT
	}{
		{getInput(test_1_tableName, test_1_columnInputArray), expectedOutput(test_1_fillableExpectedOutput, test_1_hiddenExpectedOutput)},
	}
	for _, obj := range columnTestObject {
		handler.NewColumnHandler().Handle(obj.in.tableName, obj.in.columnInputArray)
		ModelTester(t, obj.in.tableName, obj.in.columnInputArray, obj.out.expectedOutputFillable, obj.out.expectedOutputHidden)
		//MigrationTester(t, obj.in.tableName, obj.in.columnInputArray, obj.out.expectedOutputFillable, obj.out.expectedOutputHidden)
	}

}

func MigrationTester(t *testing.T, tableName string, array []models.Column, fillable []string, hidden []string) {

	migrationInfo := context.GetFromRegistry("migration").GetCtx(tableName).(*generator.MigrationGenerator).Build()
	upFunc, _ := migrationInfo.FindFunction("up")

	for _, j := range upFunc.Statements {
		fmt.Println(j)
	}

}

func ModelTester(t *testing.T, tableName string, columnArray []models.Column, fillableExpectedOutput []string, hiddenExpectedOutput []string) {

	modelClass := context.GetFromRegistry("model").GetCtx(tableName).(*generator.ModelGenerator).Build()

	var fillableData = getFillableRhs(modelClass, t)
	var hiddenData = getHiddenRhs(modelClass, t)
	arrayEquilizer(t, fillableData, fillableExpectedOutput)
	arrayEquilizer(t, hiddenData, hiddenExpectedOutput)

}

func getFillableRhs(klass *core.Class, t *testing.T) []string {
	element, err := klass.FindMember("fillable")
	if err != nil {
		t.Error("fillable not found in klass")
	}
	return (*element).(*core.ArrayAssignment).Rhs
}

func getHiddenRhs(klass *core.Class, t *testing.T) []string {
	element, err := klass.FindMember("visible")
	if err != nil {
		t.Error("hidden not found in klass")
	}
	return (*element).(*core.ArrayAssignment).Rhs
}

func arrayEquilizer(t *testing.T, in []string, out []string) {
	var table = []*struct {
		in  []string
		out []string
	}{
		{in, out},
	}

	for i, element := range table {
		for j, inner := range element.in {
			if !reflect.DeepEqual(inner, element.out[j]) {
				t.Errorf("in test case %d , %d expected '%s' found '%s'", i, j, inner, element.out[j])
			}
		}
	}
}
