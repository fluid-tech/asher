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
func TestFunctionBuilder(t *testing.T) {
	assigmentSS := core.TabbedUnit(core.GetSimpleStatement("$this->mutator = $mutator"))
	assigmentSS2:= core.TabbedUnit(core.GetSimpleStatement("$this->query = $query"))
	builder := builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArgument("BaseMutator $mutator").AddArgument("BaseQuery $query").AddStatement(&assigmentSS2).
		AddStatement(&assigmentSS)
	if builder.GetFunction().String() != Ctor{
		t.Errorf("expected \n%s \nfound:\n%s", Ctor, builder.GetFunction().String())
	}
}
