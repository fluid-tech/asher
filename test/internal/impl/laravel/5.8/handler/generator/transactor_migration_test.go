package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestTransactorMigration(t *testing.T) {
	var table = []*api.GeneralTest{
		genTransactorMigration("Teacher", TeacherMigrationForFileURLS),
		genTransactorMigration("Admin", AdminMigrationForFileURLS),
	}

	api.IterateAndTest(table, t)

}

func genTransactorMigration(className string, expectedOut string) *api.GeneralTest {
	migGen := generator.NewMigrationGenerator().SetName(className)
	generator.NewTransactorMigration(migGen).AddMigrationForFileUrls()
	return api.NewGeneralTest(migGen.String(), expectedOut)
}
