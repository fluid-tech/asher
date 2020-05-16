package builder

import (
	api2 "asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
	"asher/test/api"
	"testing"
)

func TestClassBuilder(t *testing.T) {
	var table = []*api.GeneralTest{
		getClassWithExtendsAndInitialization(),
		getClassWithoutExtendsAndInitialization(),
		buildClassBuilderWithExistingClass(),
	}

	api.IterateAndTest(table, t)
}

func getClassWithoutExtendsAndInitialization() *api.GeneralTest {
	assigmentSS := api2.TabbedUnit(core.NewSimpleStatement("$this->fullyQualifiedModel = $fullyQualifiedModel"))
	//assigmentSS2 := core.TabbedUnit(core.NewSimpleStatement("$this->query = $query"))
	functionBuilder := builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArgument("string $fullyQualifiedModel").
		AddStatement(&assigmentSS)

	member := api2.TabbedUnit(core.GetVarDeclaration("private", "fullyQualifiedModel"))

	klass := builder.NewClassBuilder().SetName("TestMutator").
		AddFunction(functionBuilder.GetFunction()).AddMember(&member).
		SetPackage("App").AddImport(`Illuminate\Database\Eloquent\Model`)

	return api.NewGeneralTest(klass.GetClass().String(), TestClass2)
}



func getClassWithExtendsAndInitialization() *api.GeneralTest {
	assigmentSS := api2.TabbedUnit(core.NewSimpleStatement("$this->fullyQualifiedModel = $fullyQualifiedModel"))
	//assigmentSS2 := core.TabbedUnit(core.NewSimpleStatement("$this->query = $query"))
	functionBuilder := builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArgument("string $fullyQualifiedModel").
		AddStatement(&assigmentSS)

	member := api2.TabbedUnit(core.GetVarDeclaration("private", "fullyQualifiedModel"))

	klass := builder.NewClassBuilder().SetName("TestMutator").SetExtends("BaseMutator").
		AddFunction(functionBuilder.GetFunction()).AddMember(&member).
		SetPackage("App").AddImport(`Illuminate\Database\Eloquent\Model`)
	return api.NewGeneralTest(klass.GetClass().String(), TestClass)
}

func buildClassBuilderWithExistingClass() *api.GeneralTest {
	assigmentSS := api2.TabbedUnit(core.NewSimpleStatement("$this->fullyQualifiedModel = $fullyQualifiedModel"))
	//assigmentSS2 := core.TabbedUnit(core.NewSimpleStatement("$this->query = $query"))
	functionBuilder := builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArgument("string $fullyQualifiedModel").
		AddStatement(&assigmentSS)

	member := api2.TabbedUnit(core.GetVarDeclaration("private", "fullyQualifiedModel"))

	klass := core.NewClass()
	klass.Name = "Hello"
	klass.Package = "Test"

	b := builder.NewClassBuilderFromClass(klass).AddFunction(functionBuilder.GetFunction()).AddMember(&member).
		SetPackage("Test")

	return api.NewGeneralTest(b.GetClass().String(), TestClass3)
}
