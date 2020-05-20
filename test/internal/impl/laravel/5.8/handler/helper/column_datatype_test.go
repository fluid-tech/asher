package helper

import (
	"asher/internal/impl/laravel/5.8/handler/helper"
	"reflect"
	"testing"
)

type IN struct {
	colType string
	colName string
	allowed []string
}

type OUT struct {
	output string
}

func GetInput(colType string, colName string, allowed []string) IN {
	return IN{colType, colName, allowed}
}

func expectedOutput(expectedOutput string) OUT {
	return OUT{expectedOutput}
}

func Test_Columns(t *testing.T) {

	var columnTestObject = []*struct {
		in  IN
		out OUT
	}{
		{GetInput("unsignedBigInteger", "desc", nil), expectedOutput("unsignedBigInteger('desc')")},
		{GetInput("bigInteger", "desc", nil), expectedOutput("bigInteger('desc')")},
		{GetInput("unsignedInteger", "desc", nil), expectedOutput("unsignedInteger('desc')")},
		{GetInput("integer", "desc", nil), expectedOutput("integer('desc')")},
		{GetInput("unsignedTinyInteger", "desc", nil), expectedOutput("unsignedTinyInteger('desc')")},
		{GetInput("tinyInteger", "desc", nil), expectedOutput("tinyInteger('desc')")},
		{GetInput("unsignedMediumInteger", "desc", nil), expectedOutput("unsignedMediumInteger('desc')")},
		{GetInput("stringsd", "desc", nil), expectedOutput("unsupported datatype")},
		{GetInput("char|12", "desc", nil), expectedOutput("char('desc', 12)")},
		{GetInput("enum", "desc", nil), expectedOutput("enum('desc')")},
		{GetInput("enum", "desc", []string{"1", "2", "3"}), expectedOutput(`enum('desc', ['1', '2', '3'])`)},

	}
	for i, obj := range columnTestObject {
		actualOutput:= helper.ColTypeSwitcher(obj.in.colType, obj.in.colName, obj.in.allowed)
		if !reflect.DeepEqual(actualOutput, obj.out.output) {
			t.Errorf("in test case %d, expected '%s' found '%s'", i, obj.out.output, actualOutput )
		}
	}




}
