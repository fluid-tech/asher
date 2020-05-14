package builder

import (
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
	"testing"
)

func TestSimpleStatementBuild(t *testing.T) {
	var table = []struct {
		in  core.SimpleStatement
		out string
	}{
		{builder.GetSimpleStatementBuilder().SetStatement(`$hello="world"`).GetStatement(),
			`$hello="world";`},
		{builder.GetSimpleStatementBuilder().SetStatement(`$hello="random"`).GetStatement(),
			`$hello="random";`},

	}

	for _, element := range table {
		if element.in.String() != element.out{
			t.Errorf("expected '%s' found '%s'", element.out, element.in.String())
		}
	}
}