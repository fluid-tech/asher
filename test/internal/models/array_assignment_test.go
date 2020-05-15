package models

import (
	"asher/internal/api/codebuilder/php/core"
	"asher/test/api"
	"testing"
)
const rhsOut =
`public $fillable = ["user_id", 
"password"
];
`
const rhsOut2 =
`private $fillable = ["user_id", 
"password"
];
`

const rhsOut3 =
` $fillable = ["user_id", 
"password"
];
`
func TestArrayAssignment(t *testing.T) {
	rhs:=[]string{`"user_id"`, `"password"`}
	var table = []*api.GeneralTest {
		api.NewGeneralTest(core.NewArrayAssignment("public", "fillable", rhs).String(), rhsOut),
		api.NewGeneralTest(core.NewArrayAssignment("private", "fillable", rhs).String(), rhsOut2),
		api.NewGeneralTest(core.NewArrayAssignment("", "fillable", rhs).String(), rhsOut3),

	}

	api.IterateAndTest(table, t)

}
