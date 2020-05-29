package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestTransactorModel(t *testing.T) {
	var table = []*api.GeneralTest{
		genTransactorModel("Teacher", TeacherModelWithFileURLS),
		genTransactorModel("Admin", AdminModelWithFileURLS),
	}
	api.IterateAndTest(table, t)
}

func genTransactorModel(className string, expectedOut string) *api.GeneralTest {
	modelGen := generator.NewModelGenerator().SetName(className)
	generator.NewTransactorModel(modelGen).
		AddFileUrlsValidationRules().
		AddFileUrlsToFillAbles()
	return api.NewGeneralTest(modelGen.String(), expectedOut)
}
