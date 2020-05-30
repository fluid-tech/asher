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
		{GetInput("char|12", "desc", nil), expectedOutput("char('desc', 12)")},
		{GetInput("enum", "desc", nil), expectedOutput("enum('desc')")},
		{GetInput("enum", "desc", []string{"1", "2", "3"}), expectedOutput(`enum('desc', ['1', '2', '3'])`)},
		{GetInput("year", "desc", nil), expectedOutput(`year('desc')`)},
		{GetInput("timeStampTz|0", "desc", nil), expectedOutput(`timestampTz('desc', 0)`)},
		{GetInput("timestamp|0", "desc", nil), expectedOutput(`timestamp('desc', 0)`)},
		{GetInput("timeTz|0", "desc", nil), expectedOutput(`timeTz('desc', 0)`)},
		{GetInput("time|0", "desc", nil), expectedOutput(`time('desc', 0)`)},
		{GetInput("text", "desc", nil), expectedOutput(`text('desc')`)},
		{GetInput("softDeleteTz|0", "desc", nil), expectedOutput(`softDeletesTz('desc', 0)`)},
		{GetInput("softDelete|0", "desc", nil), expectedOutput(`softDeletes('desc', 0)`)},
		{GetInput("polygon", "desc", nil), expectedOutput(`polygon('desc')`)},
		{GetInput("point", "desc", nil), expectedOutput(`point('desc')`)},
		{GetInput("nullableUuidMorphs", "desc", nil), expectedOutput(`nullableUuidMorphs('desc')`)},
		{GetInput("nullableMorphs", "desc", nil), expectedOutput(`nullableMorphs('desc')`)},
		{GetInput("multiPolygon", "desc", nil), expectedOutput(`multiPolygon('desc')`)},
		{GetInput("multiPoint", "desc", nil), expectedOutput(`multiPoint('desc')`)},
		{GetInput("multiLineString", "desc", nil), expectedOutput(`multiLineString('desc')`)},
		{GetInput("uuidMorphs", "desc", nil), expectedOutput(`uuidMorphs('desc')`)},
		{GetInput("morphs", "desc", nil), expectedOutput(`morphs('desc')`)},
		{GetInput("macAddress", "desc", nil), expectedOutput(`macAddress('desc')`)},
		{GetInput("longText", "desc", nil), expectedOutput(`longText('desc')`)},
		{GetInput("lineString", "desc", nil), expectedOutput(`lineString('desc')`)},
		{GetInput("jsonb", "desc", nil), expectedOutput(`jsonb('desc')`)},
		{GetInput("json", "desc", nil), expectedOutput(`json('desc')`)},
		{GetInput("ipAddress", "desc", nil), expectedOutput(`ipAddress('desc')`)},
		{GetInput("geometryCollection", "desc", nil), expectedOutput(`geometryCollection('desc')`)},
		{GetInput("geometry", "desc", nil), expectedOutput(`geometry('desc')`)},
		{GetInput("decimal", "desc", nil), expectedOutput(`decimal('desc')`)},
		{GetInput("decimal|8,2", "desc", nil), expectedOutput(`decimal('desc', 8,2)`)},
		{GetInput("dateTimeTz", "desc", nil), expectedOutput(`dateTimeTz('desc')`)},
		{GetInput("dateTime|0", "desc", nil), expectedOutput(`dateTime('desc', 0)`)},
		{GetInput("float", "desc", nil), expectedOutput(`float('desc')`)},
		{GetInput("float|8,2", "desc", nil), expectedOutput(`float('desc', 8,2)`)},
		{GetInput("double|8,2", "desc", nil), expectedOutput(`double('desc', 8,2)`)},
		{GetInput("double", "desc", nil), expectedOutput(`double('desc')`)},
		{GetInput("date", "desc", nil), expectedOutput(`date('desc')`)},
		{GetInput("char|100", "desc", nil), expectedOutput(`char('desc', 100)`)},
		{GetInput("boolean", "desc", nil), expectedOutput(`boolean('desc')`)},
		{GetInput("string", "desc", nil), expectedOutput(`string('desc')`)},
		{GetInput("string|100", "desc", nil), expectedOutput(`string('desc', 100)`)},
		{GetInput("mediumInteger", "desc", nil), expectedOutput(`mediumInteger('desc')`)},
		{GetInput("unsignedMediumInteger", "desc", nil), expectedOutput(`unsignedMediumInteger('desc')`)},
		{GetInput("tinyInteger", "desc", nil), expectedOutput(`tinyInteger('desc')`)},
		{GetInput("unsignedTinyInteger", "desc", nil), expectedOutput(`unsignedTinyInteger('desc')`)},
		{GetInput("integer", "desc", nil), expectedOutput(`integer('desc')`)},
		{GetInput("unsignedInteger", "desc", nil), expectedOutput(`unsignedInteger('desc')`)},
		{GetInput("bigInteger", "desc", nil), expectedOutput(`bigInteger('desc')`)},
		{GetInput("unsignedBigInteger", "desc", nil), expectedOutput(`unsignedBigInteger('desc')`)},
	}
	for i, obj := range columnTestObject {
		actualOutput, _ := helper.ColTypeSwitcher(obj.in.colType, obj.in.colName, obj.in.allowed)
		if !reflect.DeepEqual(actualOutput, obj.out.output) {
			t.Errorf("in test case %d, expected '%s' found '%s'", i, obj.out.output, actualOutput)
		}
	}

}
