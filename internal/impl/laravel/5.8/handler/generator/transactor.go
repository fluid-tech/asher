package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
)

type TransactorGenerator struct {
	classBuilder   interfaces.Class
	identifier     string
	imports        []string
	transactorType string
}

func NewTransactorGenerator(identifier string, transactorType string) *TransactorGenerator {
	return &TransactorGenerator{
		classBuilder:   builder.NewClassBuilder(),
		identifier:     identifier,
		imports:        []string{},
		transactorType: transactorType,
	}
}

/**
Sets the identifier of the current class
Parameters:
	- identifier: string
Sample Usage:
	- SetIdentifier("ClassName")
*/
func (transactorGenerator *TransactorGenerator) SetIdentifier(identifier string) *TransactorGenerator {
	transactorGenerator.identifier = identifier
	return transactorGenerator
}

/**
Sets the type of the transactor
Parameters:
	- identifier: string
Sample Usage:
	- SetTransactorType("default") or SetTransactorType("file") or SetTransactorType("image")
*/
func (transactorGenerator *TransactorGenerator) SetTransactorType(identifier string) *TransactorGenerator {
	transactorGenerator.transactorType = identifier
	return transactorGenerator
}

/**
Appends import to the controller file
Parameters:
	- units: string array of the import
Returns:
	- instance of ControllerGenerator object
Sample Usage:
	- AppendImport([]string{"App\User",})
*/
func (transactorGenerator *TransactorGenerator) AppendImports(units []string) *TransactorGenerator {
	transactorGenerator.imports = append(transactorGenerator.imports, units...)
	return transactorGenerator
}

/**
Adds Constructor in the Transactor with Query and Mutator Injected of the currentModel
Returns:
	- Return instance of TransactorGenerator
Sample Usage:
	- transactorGeneratorObject.AddConstructor()
*/
func (transactorGenerator *TransactorGenerator) AddConstructorFunction() *TransactorGenerator {
	lowerCamelIdentifier := strcase.ToLowerCamel(transactorGenerator.identifier)
	queryVariableName := lowerCamelIdentifier + `Query`
	mutatorVariableName := lowerCamelIdentifier + `Mutator`

	constructorArguments := []string{
		transactorGenerator.identifier + `Query $` + queryVariableName,
		transactorGenerator.identifier + `Mutator $` + mutatorVariableName,
	}

	parentConstructorCall := core.NewFunctionCall("parent::__construct").
		AddArg(core.NewParameter(fmt.Sprintf("$%s" , queryVariableName))).
		AddArg(core.NewParameter(fmt.Sprintf("$%s" , mutatorVariableName))).
		AddArg(core.NewParameter(`"id"`))

	switch transactorGenerator.transactorType {
	case "file":
		transactorGenerator.imports = append(transactorGenerator.imports, `use App\Helpers\BaseFileUploadHelper`)
		parentConstructorCall.AddArg(core.NewParameter(
			`new BaseFileUploadHelper(self::BASE_PATH, self::IMAGE_VALIDATION_RULES,"png")`))
		/*TODO add something for const*/
		transactorGenerator.classBuilder.AddMember(core.NewSimpleStatement(
			"public const IMAGE_VALIDATION_RULES =" +
				" array(\n        'file' => 'required|mimes:jpeg,jpg,png|max:3000'\n    )"))

		transactorGenerator.classBuilder.AddMember(
			core.NewSimpleStatement(`private const BASE_PATH = "`+
				strings.ToLower(transactorGenerator.identifier)+`"`))

	case "image":
		transactorGenerator.imports = append(transactorGenerator.imports, `use App\Helpers\ImageUploadHelper`)
		parentConstructorCall.AddArg(core.NewParameter(
			`new ImageUploadHelper(self::BASE_PATH, self::IMAGE_VALIDATION_RULES)`))
		/*TODO add something for const*/
		transactorGenerator.classBuilder.AddMember(core.NewSimpleStatement(
"public const IMAGE_VALIDATION_RULES = " +
	"array(\n        'file' => 'required|mimes:jpeg,jpg,png|max:3000'\n    )"))

		transactorGenerator.classBuilder.AddMember(
			core.NewSimpleStatement(`private const BASE_PATH = "`+
				strings.ToLower(transactorGenerator.identifier)+`"`))
	}

	constructorStatements := []api.TabbedUnit{
		parentConstructorCall,
		core.NewSimpleStatement("$this->className = self::CLASS_NAME"),
	}

	transactorGenerator.classBuilder.AddFunction(builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArguments(constructorArguments).AddStatements(constructorStatements).GetFunction())
	return transactorGenerator
}

/**
Main Function To be called when we want to build the transactor
Returns:
	- Return instance of core.Class
Sample Usage:
	- transactorGeneratorObject.BuildRestTransactor()
*/
func (transactorGenerator *TransactorGenerator) BuildTransactor() *core.Class {
	const namespace = `App\Transactors`
	var extendsTransactor string

	switch transactorGenerator.transactorType {
	case "default":
		extendsTransactor = "BaseTransactor"
	case "file":
		extendsTransactor = "FileTransactor"
	case "image":
		extendsTransactor = "ImageTransactor"
	default:
		extendsTransactor = "BaseTransactor"
	}

	transactorImports := []string{
		`App\Query\` + transactorGenerator.identifier + `Query`,
		`App\Transactors\Mutations\` + transactorGenerator.identifier + `Mutator`,
	}

	className := transactorGenerator.identifier + "Transactor"

	transactorGenerator.AppendImports(transactorImports)
	transactorGenerator.AddConstructorFunction()

	transactorGenerator.classBuilder.AddMember(core.NewSimpleStatement(
		"private const CLASS_NAME = '" + className + "'")).SetName(className).
		SetExtends(extendsTransactor).SetPackage(namespace).AddImports(transactorGenerator.imports)

	return transactorGenerator.classBuilder.GetClass()
}

/**
Returns:
	- Return string object of TransactorGenerator
Sample Usage:
	- transactorGeneratorObject.String()
*/
func (transactorGenerator *TransactorGenerator) String() string {
	return transactorGenerator.BuildTransactor().String()
}
