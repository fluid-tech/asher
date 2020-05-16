package generator

import (
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestMigrationGenerator(t *testing.T) {
	var table = []*api.GeneralTest{
		getEmptyMigrationWithName(),
		getMigrationWithColumns(),
	}
	//fmt.Println(table)
	api.IterateAndTest(table, t)
}

func getEmptyMigrationWithName() *api.GeneralTest {
	migrationGenerator := generator.NewMigrationGenerator().SetName("student_allotments")

	return api.NewGeneralTest(migrationGenerator.Build().String(), EmptyMigrationWithName)
}

func getMigrationWithColumns() *api.GeneralTest {
	columns := []core.SimpleStatement {
		*core.NewSimpleStatement("$this->string('name')"),
		*core.NewSimpleStatement("$this->string('phone_number', 12)->unique()"),
	}
	migrationGenerator := generator.NewMigrationGenerator().SetName("student_allotments").
		AddColumns(columns)

	return api.NewGeneralTest(migrationGenerator.Build().String(), MigrationWithColumns)
}
