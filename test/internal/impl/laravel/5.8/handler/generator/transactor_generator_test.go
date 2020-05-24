package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestTransactorGenerator(t *testing.T) {
	var table = []*api.GeneralTest{
		genTransactorTest("Student", "default", StudentBasicTransactor),
		genTransactorTest("Admin", "file", AdminFileTransactor),
		genTransactorTest("Teacher", "image", TeacherImageTransactor),
	}
	api.IterateAndTest(table, t)
}

func genTransactorTest(modelName string, transactorType string, expectedOut string) *api.GeneralTest {
	transactorGenerator := generator.NewTransactorGenerator(modelName, transactorType)
	//fmt.Print(transactorGenerator.String())
	return api.NewGeneralTest(transactorGenerator.String(), expectedOut)
}
