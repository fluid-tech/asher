package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestQueryGenerator(t *testing.T) {
	//transactorGenerator := generator.NewQueryGenerator("Order", true)
	//fmt.Print(transactorGenerator)
	var table = []*api.GeneralTest{
		genQueryTest("Student", StudentBasicQuery),
		genQueryTest("Admin", AdminBasicQuery),
		genQueryTest("Teacher", TeacherBasicQuery),
	}
	api.IterateAndTest(table, t)
}

func genQueryTest(modelName string, expectedOut string) *api.GeneralTest {
	/*TODO relation is not used for iteration1*/
	transactorGenerator := generator.NewQueryGenerator(true).SetIdentifier(modelName)
	return api.NewGeneralTest(transactorGenerator.String(), expectedOut)
}
