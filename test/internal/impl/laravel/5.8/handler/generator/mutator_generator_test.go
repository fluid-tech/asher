package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestMutatorGenerator(t *testing.T) {

	var table = []*api.GeneralTest{
		genMutatorTest("Student", StudentBasicMutator),
		genMutatorTest("Admin", AdminBasicMutator),
		genMutatorTest("Teacher", TeacherBasicMutator),
	}
	api.IterateAndTest(table, t)
}

func genMutatorTest(modelName string, expectedOut string) *api.GeneralTest {
	/*TODO relation is not used for iteration1*/
	mutatorgen := generator.NewMutatorGenerator()
	mutatorgen.SetIdentifier(modelName)
	return api.NewGeneralTest(mutatorgen.String(), expectedOut)
}
