package builder

import (
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
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
	assigmentSS := core.TabbedUnit(core.GetSimpleStatement("$this->mutator = $mutator"))
	assigmentSS2:= core.TabbedUnit(core.GetSimpleStatement("$this->query = $query"))
	builder := builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArgument("BaseMutator $mutator").AddArgument("BaseQuery $query").AddStatement(&assigmentSS2).
		AddStatement(&assigmentSS)

	var table = []struct{
		in string
		out string
	}{
		{builder.GetFunction().String(), Ctor},
		{builder.AddArgument("ImageHandler $imageHandler").GetFunction().String(), Ctor2},
	}

	for _, element := range table{
		if element.in != element.out {
			t.Errorf("expected '%s' found '%s'", element.in, element.out)
		}
	}
}
