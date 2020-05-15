package builder

import (
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
	"fmt"
	"testing"
)
const TEST_CLASS string =
`namespace App;

use Illuminate\Database\Eloquent\Model;

class JobProfile extends Model
{
    public $table = "op_job_profiles";
}
`
func TestClassBuilder(t *testing.T)  {

	assignmentSS := &core.SimpleStatement{
		SimpleStatement: "$this->mutator = mutator",
	}
	function := &builder.Function{}
	function.SetName("hello").SetVisibility("public")
	function.AddArgument("BaseMutator mutator")
	tab := core.TabbedUnit(assignmentSS)
	function.AddStatement(&tab)

	fmt.Println(function.GetFunction().String())

	statement := &core.SimpleStatement{
		SimpleStatement: `public $table="op_job_profiles"`,
	}
	klass := builder.Class{}
	klass.SetName("JobProfile").SetExtends("Model")
	member := core.TabbedUnit(statement)
	klass.AddMembers([]*core.TabbedUnit{
		&member,
	}).AddFunction(function.GetFunction())

	fmt.Println(len(klass.GetClass().Members))

	if klass.GetClass().String() != TEST_CLASS{
		t.Error("expected ", klass.GetClass().String(), " \n\n\n found", TEST_CLASS)
	}

}
