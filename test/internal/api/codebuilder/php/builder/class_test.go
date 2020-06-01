package builder

import (
	api2 "asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"asher/test/api"
	"testing"
)

const (
	FullyQualifiedModelAssignment = "$this->fullyQualifiedModel = $fullyQualifiedModel"
	FullyQualifiedModelArgs       = `string $fullyQualifiedModel`
)

func TestClassBuilder(t *testing.T) {
	var table = []*api.GeneralTest{
		genClassTest("TestMutator", "App", "BaseMutator", "", `Illuminate\Database\Eloquent\Model`, TestClass),
		genClassTest("TestMutator", "App", "", "", `Illuminate\Database\Eloquent\Model`, TestClass2),
		genClassTest("Hello", "Test", "", "", "", TestClass3),
		genClassTest("Hello", "Test", "", "Runnable", "", TestClass4),
	}

	api.IterateAndTest(table, t)
}

/*
Generates Test Cases.
TODO: Replace the input string from an Object of Input Class for better future management of Inputs.
*/
func genClassTest(className string, packageName string, extendedClassName string, implementedInterfaceName string, importStatement string, expectedOutput string) *api.GeneralTest {
	functionBuilder := constructorFunction()
	member := api2.TabbedUnit(core.NewVarDeclaration("private", "fullyQualifiedModel"))

	klass := core.NewClass()
	klass.Name = className
	klass.Package = packageName

	b := builder.NewClassBuilderFromClass(klass).AddFunction(functionBuilder.GetFunction()).AddMember(member).
		SetPackage(packageName)

	if importStatement != "" {
		b.AddImport(importStatement)
	}

	if extendedClassName != "" {
		b.SetExtends(extendedClassName)
	}

	if implementedInterfaceName != "" {
		b.AddInterface(implementedInterfaceName)
	}
	return api.NewGeneralTest(b.GetClass().String(), expectedOutput)
}

/*
Generates Constructor:
Output:
	`public function __construct(string $fullyQualifiedModel) {
        $this->fullyQualifiedModel = $fullyQualifiedModel;
    }`
*/
func constructorFunction() interfaces.Function {
	assigmentSS := core.NewSimpleStatement(FullyQualifiedModelAssignment)
	return builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArgument(FullyQualifiedModelArgs).
		AddStatement(assigmentSS)
}
