package handler

import (
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler"
	"asher/internal/impl/laravel/5.8/handler/context"
	"reflect"
	"testing"
)
var arr = []string{`"user_id"`, `"password"`, `"created_by"`, `"updated_by"`, `"created_at"`, `"updated_at"`, `"deleted_at"`}
var brr = []string{`"user_id"`, `"password"`, `"created_by"`, `"updated_by"`, `"deleted_at"`}
var crr = []string{`"user_id"`, `"password"`, `"created_by"`, `"updated_by"`}
var drr = []string{`"user_id"`, `"password"`, `"created_at"`, `"updated_at"`, `"deleted_at"`}
var frr = []string{`"user_id"`, `"password"`}
func TestAuditCol(t *testing.T){
	handleFillableTestCases(t)
}

func handleFillableTestCases(t *testing.T) {
	var table = []*struct{
		in  []string
		out []string
	}{
		{genCoreAssignmentRhs("Hello", true, true, true, t), arr},
		{genCoreAssignmentRhs("HelloWorld", true, true, false, t), brr},
		{genCoreAssignmentRhs("HelloRandom", true, false, false, t), crr},
		{genCoreAssignmentRhs("HelloWorld2", false, true, true, t), drr},
		{genCoreAssignmentRhs("HelloRandom", false, false, false, t), frr},
	}

	for i, element := range table {
		if !reflect.DeepEqual(element.in, element.out) {
			t.Errorf("in test case %d expected '%s' found '%s'", i,  element.in, element.out)
		}
	}
}

func genCoreAssignmentRhs(className string, auditCol bool, softDeletes bool, timestamp bool, t *testing.T) []string {
	klass := buildClassWithArrayDecl(className)
	// adding klass to model registry
	modelRegistry := context.GetFromRegistry("model")
	modelRegistry.AddToCtx(className, klass)

	// calling handle for fillable
	auditColHandler := handler.NewAuditColHandler()
	auditColHandler.Handle(className, handler.NewAuditColInputFromType(auditCol, softDeletes, timestamp, className))

	element, err := klass.FindInMembers("fillable")
	if err != nil {
		t.Error("fillable not found in klass")
	}
	return (*element).(*core.ArrayAssignment).Rhs
}

func buildClassWithArrayDecl(className string) *core.Class{
	assigmentSS := core.TabbedUnit(core.NewSimpleStatement("$this->fullyQualifiedModel = $fullyQualifiedModel"))
	functionBuilder := builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArgument("string $fullyQualifiedModel").
		AddStatement(&assigmentSS)

	member := core.TabbedUnit(core.GetVarDeclaration("private", "fullyQualifiedModel"))
	rhs:=[]string{`"user_id"`, `"password"`}
	arrayAssignmentMember := core.TabbedUnit(core.NewArrayAssignment("public", "fillable", rhs))
	klass := builder.NewClassBuilder().SetName(className).SetExtends("BaseMutator").
		AddFunction(functionBuilder.GetFunction()).AddMember(&member).AddMember(&arrayAssignmentMember).
		SetPackage("App").AddImport(`Illuminate\Database\Eloquent\Model`)
	return klass.GetClass()
}