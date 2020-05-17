package generator

import (
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestMigrationGenerator(t *testing.T) {
	var table = []*api.GeneralTest{
		createMigrationGeneratorTest("student_allotments", []core.SimpleStatement{}, EmptyMigrationWithName),
		createMigrationGeneratorTest("student_allotments", []core.SimpleStatement {
			*core.NewSimpleStatement("$this->string('name')"),
			*core.NewSimpleStatement("$this->string('phone_number', 12)->unique()"),
		}, MigrationWithColumns),
	}
	//fmt.Println(table)
	api.IterateAndTest(table, t)
}

/**
 An internal function to create general test cases for MigrationGenerator
 */
func createMigrationGeneratorTest(name string, columns []core.SimpleStatement, expectedCode string) *api.GeneralTest {
	migrationGenerator := generator.NewMigrationGenerator().SetName(name).
		AddColumns(columns)

	return api.NewGeneralTest(migrationGenerator.Build().String(), expectedCode)
}
