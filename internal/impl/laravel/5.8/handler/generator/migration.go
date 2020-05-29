package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
)

const (
	MigrationExtends = "Migration"
	IdentifierDown = "down"
	IdentifierUp = "up"
)

type MigrationGenerator struct {
	api.Generator
	classBuilder interfaces.Class
	tableName    string
	columns      []*core.SimpleStatement
}

/**
Creates a new instance of this generator with a new interfaces.Class
*/
func NewMigrationGenerator() *MigrationGenerator {
	return &MigrationGenerator{
		classBuilder: builder.NewClassBuilder(),
		columns:      []*core.SimpleStatement{},
	}
}

/**
 Set's the name of the migration class.
 Parameters:
	- tableName: the name of table specified in asher config.
 Returns:
	- instance of the migration generator object.
 Example:
	- SetName('product_categories')
*/
func (migrationGenerator *MigrationGenerator) SetName(tableName string) *MigrationGenerator {
	className := "Create" + strcase.ToCamel(tableName) + "Table"
	migrationGenerator.classBuilder.SetName(className)
	migrationGenerator.tableName = strcase.ToSnake(tableName)
	return migrationGenerator
}

/**
 Adds the given column to the up function of this migration.
 Parameters:
	- column: core.SimpleStatement of the column to add.
 Returns:
	- instance of the migration generator object.
 Example:
	- AddColumn(core.NewSimpleStatement(`$this->string("name")->unique()`))
*/
func (migrationGenerator *MigrationGenerator) AddColumn(column *core.SimpleStatement) *MigrationGenerator {
	return migrationGenerator.AddColumns([]*core.SimpleStatement{column})
}

/**
 Adds the given list of columns to the up function of this migration.
 Parameters:
	- columns: A list of core.SimpleStatement of the columns to add.
 Returns:
	- instance of the migration generator object.
 Example:
	- AddColumns([]core.SimpleStatement{
		core.NewSimpleStatement(`$this->string('name')->unique()`),
		core.NewSimpleStatement(`$this->string('phone_number', 12)->unique()`)
	  })
*/
func (migrationGenerator *MigrationGenerator) AddColumns(columns []*core.SimpleStatement) *MigrationGenerator {
	migrationGenerator.columns = append(migrationGenerator.columns, columns...)
	return migrationGenerator
}

/**
 Generates corresponding core.Class from the given parameters of this migration.
 Returns:
	- a pointer to the corresponding core.Class of this migration
*/
func (migrationGenerator *MigrationGenerator) Build() *core.Class {
	// Preparing the arguments for up function

	arg1 := core.NewParameter("'" + migrationGenerator.tableName + "'")
	closure := builder.NewFunctionBuilder().AddArgument("Blueprint $table")
	for _, element := range migrationGenerator.columns {
		closure.AddStatement(element)
	}

	arg2 := closure.GetFunction()

	// Preparing the statements for up function

	schemaBlock := core.NewFunctionCall("Schema::create").AddArg(arg1).AddArg(arg2)

	upFunction := builder.NewFunctionBuilder().SetName(IdentifierUp).SetVisibility(VisibilityPublic).
		AddStatement(schemaBlock).GetFunction()

	// Preparing the statements for down function
	dropStatement := api.TabbedUnit(core.NewSimpleStatement("Schema::dropIfExists('" + migrationGenerator.tableName + "')"))
	downFunction := builder.NewFunctionBuilder().SetName(IdentifierDown).SetVisibility(VisibilityPublic).
		AddStatement(dropStatement).GetFunction()

	migrationGenerator.classBuilder.AddImports([]string{
		`Illuminate\Database\Migrations\Migration`,
		`Illuminate\DatabaseSchema\Blueprint`,
		`Illuminate\Support\Facades\Schema`,
	}).SetExtends(MigrationExtends).AddFunction(upFunction).
		AddFunction(downFunction)

	return migrationGenerator.classBuilder.GetClass()
}

/**
 Implementation of the base Generator to return this migration's corresponding string.
 Returns:
	- string representation of this migration.
*/
func (migrationGenerator *MigrationGenerator) String() string {
	return migrationGenerator.Build().String()
}
