package helper

import (
	"asher/internal/impl/laravel/5.8/handler/helper"
	"asher/test/api"
	"testing"
)

func Test_Column_Datatype(t *testing.T) {
	ColumnTester(t)
	PrimaryColumnTester(t)
}

func ColumnTester(t *testing.T) {

	var columnTestObject = []*api.GeneralTest{
		columnTester("uasdasdasdnsignedBigIntegeraaaa", "desc", nil, ""),
		columnTester("unsignedBigInteger", "desc", nil, "unsignedBigInteger('desc')"),
		columnTester("bigInteger", "desc", nil, "bigInteger('desc')"),
		columnTester("unsignedInteger", "desc", nil, "unsignedInteger('desc')"),
		columnTester("integer", "desc", nil, "integer('desc')"),
		columnTester("unsignedTinyInteger", "desc", nil, "unsignedTinyInteger('desc')"),
		columnTester("tinyInteger", "desc", nil, "tinyInteger('desc')"),
		columnTester("unsignedMediumInteger", "desc", nil, "unsignedMediumInteger('desc')"),
		columnTester("char|12", "desc", nil, "char('desc', 12)"),
		columnTester("enum", "desc", nil, "enum('desc')"),
		columnTester("enum", "desc", []string{"1", "2", "3"}, "enum('desc', ['1', '2', '3'])"),
		columnTester("set", "desc", nil, "set('desc')"),
		columnTester("set", "desc", []string{"1", "2", "3"}, "set('desc', ['1', '2', '3'])"),
		columnTester("year", "desc", nil, "year('desc')"),
		columnTester("timeStampTz|0", "desc", nil, "timestampTz('desc', 0)"),
		columnTester("timestamp|0", "desc", nil, "timestamp('desc', 0)"),
		columnTester("timeTz|0", "desc", nil, "timeTz('desc', 0)"),
		columnTester("time|0", "desc", nil, "time('desc', 0)"),
		columnTester("text", "desc", nil, "text('desc')"),
		columnTester("softDeleteTz|0", "desc", nil, "softDeletesTz('desc', 0)"),
		columnTester("softDelete|0", "desc", nil, "softDeletes('desc', 0)"),
		columnTester("polygon", "desc", nil, "polygon('desc')"),
		columnTester("point", "desc", nil, "point('desc')"),
		columnTester("nullableUuidMorphs", "desc", nil, "nullableUuidMorphs('desc')"),
		columnTester("nullableMorphs", "desc", nil, "nullableMorphs('desc')"),
		columnTester("multiPolygon", "desc", nil, "multiPolygon('desc')"),
		columnTester("multiPoint", "desc", nil, "multiPoint('desc')"),
		columnTester("multiLineString", "desc", nil, "multiLineString('desc')"),
		columnTester("uuidMorphs", "desc", nil, "uuidMorphs('desc')"),
		columnTester("morphs", "desc", nil, "morphs('desc')"),
		columnTester("macAddress", "desc", nil, "macAddress('desc')"),
		columnTester("longText", "desc", nil, "longText('desc')"),
		columnTester("lineString", "desc", nil, "lineString('desc')"),
		columnTester("jsonb", "desc", nil, "jsonb('desc')"),
		columnTester("json", "desc", nil, "json('desc')"),
		columnTester("ipAddress", "desc", nil, "ipAddress('desc')"),
		columnTester("geometryCollection", "desc", nil, "geometryCollection('desc')"),
		columnTester("geometry", "desc", nil, "geometry('desc')"),
		columnTester("decimal", "desc", nil, "decimal('desc')"),
		columnTester("decimal|8,2", "desc", nil, "decimal('desc', 8,2)"),
		columnTester("dateTimeTz", "desc", nil, "dateTimeTz('desc')"),
		columnTester("dateTime|0", "desc", nil, "dateTime('desc', 0)"),
		columnTester("float", "desc", nil, "float('desc')"),
		columnTester("float|8,2", "desc", nil, "float('desc', 8,2)"),
		columnTester("double|8,2", "desc", nil, "double('desc', 8,2)"),
		columnTester("double", "desc", nil, "double('desc')"),
		columnTester("date", "desc", nil, "date('desc')"),
		columnTester("char|100", "desc", nil, "char('desc', 100)"),
		columnTester("boolean", "desc", nil, "boolean('desc')"),
		columnTester("string", "desc", nil, "string('desc')"),
		columnTester("string|100", "desc", nil, "string('desc', 100)"),
		columnTester("mediumInteger", "desc", nil, "mediumInteger('desc')"),
		columnTester("unsignedMediumInteger", "desc", nil, "unsignedMediumInteger('desc')"),
		columnTester("tinyInteger", "desc", nil, "tinyInteger('desc')"),
		columnTester("unsignedTinyInteger", "desc", nil, "unsignedTinyInteger('desc')"),
		columnTester("integer", "desc", nil, "integer('desc')"),
		columnTester("unsignedInteger", "desc", nil, "unsignedInteger('desc')"),
		columnTester("bigInteger", "desc", nil, "bigInteger('desc')"),
		columnTester("unsignedBigInteger", "desc", nil, "unsignedBigInteger('desc')"),
		columnTester("smallInteger", "desc", nil, "smallInteger('desc')"),
		columnTester("unsignedSmallInteger", "desc", nil, "unsignedSmallInteger('desc')"),
		columnTester("binary", "desc", nil, "binary('desc')"),
	}
	api.IterateAndTest(columnTestObject, t)

}

func columnTester(colType string, colName string, allowed []string, expectedOutput string) *api.GeneralTest {
	actualOutput, _ := helper.ColTypeSwitcher(colType, colName, allowed)
	return api.NewGeneralTest(actualOutput, expectedOutput)
}

func PrimaryColumnTester(t *testing.T) {
	var table = []*api.GeneralTest{
		genPrimaryKeyTest("unknown Type", ""),
		genPrimaryKeyTest("integer", "increments"),
		genPrimaryKeyTest("bigInteger", "bigIncrements"),
		genPrimaryKeyTest("tinyInteger", "tinyIncrements"),
		genPrimaryKeyTest("smallInteger", "smallIncrements"),
		genPrimaryKeyTest("mediumInteger", "mediumIncrements"),
	}
	api.IterateAndTest(table, t)
}

func genPrimaryKeyTest(key string, expectedOut string) *api.GeneralTest {
	generated, _ := helper.PrimaryKeyMethodNameGenerator(key)
	return api.NewGeneralTest(generated, expectedOut)
}
