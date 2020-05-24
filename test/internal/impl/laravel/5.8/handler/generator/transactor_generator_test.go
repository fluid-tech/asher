package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestTransactorGenerator(t *testing.T) {
	var table = []*api.GeneralTest{
		genTransactorTest("Order", "default", BasicTransactor),
		genTransactorTest("Order", "file", FileTransactor),
		genTransactorTest("Order", "image", ImageTransactor),
		genTransactorTest("Order", "", BasicTransactor),
	}
	api.IterateAndTest(table, t)
}

func genTransactorTest(modelName string, transactorType string, expectedOut string) *api.GeneralTest {
	transactorGenerator := generator.NewTransactorGenerator(modelName, transactorType)

	return api.NewGeneralTest(transactorGenerator.String(), expectedOut)
}
