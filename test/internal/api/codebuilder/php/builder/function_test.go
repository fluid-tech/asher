package builder

import (
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
	"asher/test/api"
	"testing"
)

const Ctor =
`public function __construct(BaseMutator $mutator, BaseQuery $query) {
    $this->query = $query;
    $this->mutator = $mutator;
}

`
const Ctor2 =
`public function __construct(BaseMutator $mutator, BaseQuery $query, ImageHandler $imageHandler) {
    $this->query = $query;
    $this->mutator = $mutator;
}

`
func TestFunctionBuilder(t *testing.T) {
	assigmentSS := core.TabbedUnit(core.NewSimpleStatement("$this->mutator = $mutator"))
	assigmentSS2:= core.TabbedUnit(core.NewSimpleStatement("$this->query = $query"))
	builder := builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArgument("BaseMutator $mutator").AddArgument("BaseQuery $query").AddStatement(&assigmentSS2).
		AddStatement(&assigmentSS)

	var table = []*api.GeneralTest{
		api.NewGeneralTest(builder.GetFunction().String(), Ctor),
		api.NewGeneralTest(builder.AddArgument("ImageHandler $imageHandler").GetFunction().String(), Ctor2),
		buildFunctionBuilderWithExistingFunction(),
	}

	api.IterateAndTest(table, t)
}

func buildFunctionBuilderWithExistingFunction() *api.GeneralTest {
	function := core.NewFunction()
	function.Name = "up"
	function.Visibility = "protected"
	function.Arguments = []string{"$hello", "$world"}
	s := core.TabbedUnit(core.NewSimpleStatement("return $world+$hello"))
	b := builder.NewFunctionBuilderFromFunction(function).AddStatement(&s)
	return api.NewGeneralTest(b.GetFunction().String(), TestFunction)
}


const TestFunction=`protected function up($hello, $world) {
    return $world+$hello;
}

`