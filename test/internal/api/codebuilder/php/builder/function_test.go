package builder

import (
	api2 "asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
	"asher/test/api"
	"testing"
)

func TestFunctionBuilder(t *testing.T) {

	var table = []*api.GeneralTest{

		genFuncTest("__construct", "public", true, []string{"BaseMutator $mutator", "BaseQuery $query"}, []api2.TabbedUnit{
			core.NewSimpleStatement("$this->query = $query"),
			core.NewSimpleStatement("$this->mutator = $mutator"),
		}, Ctor),

		genFuncTest("__construct", "public", true, []string{"BaseMutator $mutator", "BaseQuery $query", "ImageHandler $imageHandler"}, []api2.TabbedUnit{
			core.NewSimpleStatement("$this->query = $query"),
			core.NewSimpleStatement("$this->mutator = $mutator"),
		}, Ctor2),

		genFuncTest("up", "protected", false, []string{"$hello", "$world"}, []api2.TabbedUnit{
			core.NewSimpleStatement("return $world+$hello"),
		}, TestFunction),
	}

	api.IterateAndTest(table, t)
}

func genFuncTest(funcName string, funcVisibility string, isStatic bool, args []string, statements []api2.TabbedUnit, expectedOutput string) *api.GeneralTest {
	function := core.NewFunction()
	function.Name = funcName
	function.Visibility = funcVisibility
	function.Arguments = args
	b := builder.NewFunctionBuilderFromFunction(function).AddStatements(statements).SetStatic(isStatic)
	return api.NewGeneralTest(b.GetFunction().String(), expectedOutput)
}
