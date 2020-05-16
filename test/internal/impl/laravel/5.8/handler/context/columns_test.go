package context

import (
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/models"
	"reflect"
	"testing"
)

func Test_Columns(t *testing.T) {

	var columnTestObject = []*struct {
		tableName              string
		columnInputArray       []models.Column
		expectedOutputFillable []string
		expectedOutputHidden   []string
	}{
		{test_1_tableName, test_1_columnInputArray, test_1_fillableExpectedOutput, test_1_hiddenExpectedOutput},
		{test_1_tableName, test_1_columnInputArray, test_1_fillableExpectedOutput, test_1_hiddenExpectedOutput},
	}
	for _, obj := range columnTestObject {
		ModelTester(t, obj.tableName, obj.columnInputArray, obj.expectedOutputFillable, obj.expectedOutputHidden)
	}

}

func ModelTester(t *testing.T, tableName string, columnArray []models.Column, fillableExpectedOutput []string, hiddenExpectedOutput []string) {

	handler.NewColumnHandler().Handle(tableName, columnArray)

	modelClass := context.GetFromRegistry("model").GetCtx(tableName).(*core.Class)

	var fillableData []string = getFillableRhs(modelClass, t)
	var hiddenData []string = getHiddenRhs(modelClass, t)
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
