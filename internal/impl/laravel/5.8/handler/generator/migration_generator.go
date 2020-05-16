package generator

import (
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
)

type MigrationGenerator struct {
	classBuilder interfaces.Class
	tableName string
	columns []core.SimpleStatement
}

/**
 Creates a new instance of this generator with a new interfaces.Class
*/
func NewMigrationGenerator() *MigrationGenerator {
	return &MigrationGenerator{
		classBuilder: builder.NewClassBuilder(),
	}
}

/**
 Set's the name of the migration class.
 Parameters:
	- tableName: the name of table specified in asher config.
 Returns:
	- instance of the migration generator object.
*/
func (migrationGenerator *MigrationGenerator) SetName(columnName string) *MigrationGenerator {
	className := "Create" + strcase.ToCamel(columnName) + "Table"
	migrationGenerator.classBuilder.SetName(className)
	return migrationGenerator
}

/**
 Adds the given column to the up function of this migration.
 Parameters:
	- column: core.SimpleStatement of the column to add.
 Returns:
	- instance of the migration generator object.
 */
func (migrationGenerator *MigrationGenerator) AddColumn(column core.SimpleStatement) *MigrationGenerator {
	migrationGenerator.columns = append(migrationGenerator.columns, column)
	return migrationGenerator
}

/**
 Adds the given list of columns to the up function of this migration.
 Parameters:
	- columns: A list of core.SimpleStatement of the columns to add.
 Returns:
	- instance of the migration generator object.
*/
func (migrationGenerator *MigrationGenerator) AddColumns(columns []core.SimpleStatement) *MigrationGenerator {
	for _, statement := range columns {
		migrationGenerator.AddColumn(statement)
	}
	return migrationGenerator
}

/**
 Generates corresponding core.Class from the given parameters of this migration.
 Returns:
	- a pointer to the corresponding core.Class of this migration
 */
func (migrationGenerator *MigrationGenerator) Build() *core.Class {
	// Preparing the arguments for up function
	arg1 := core.TabbedUnit(core.NewArgument(migrationGenerator.tableName))
	closure := builder.NewFunctionBuilder().AddArgument("Blueprint $table")
	for _, stmt := range migrationGenerator.columns {
		tabbedStatement := core.TabbedUnit(&stmt)
		closure.AddStatement(&tabbedStatement)
	}
	arg2 := core.TabbedUnit(closure.GetFunction())

	// Preparing the statements for up function
	schemaBlock := core.TabbedUnit(core.NewFunctionCall("Schema::create").AddArg(&arg1).AddArg(&arg2))
	upFunction := builder.NewFunctionBuilder().SetName("up").SetVisibility("public").AddStatement(&schemaBlock).GetFunction()

	// Preparing the statements for down function
	dropStatement := core.TabbedUnit(core.NewSimpleStatement("Schema::dropIfExists(" + migrationGenerator.tableName + ")"))
	downFunction := builder.NewFunctionBuilder().SetName("down").SetVisibility("public").AddStatement(&dropStatement).GetFunction()

	migrationGenerator.classBuilder.AddImports([]string{
		`Illuminate\Database\Migrations\Migration`,
		`Illuminate\DatabaseSchema\Blueprint`,
		`Illuminate\Support\Facades\Schema`,
	}).SetExtends("Migration").AddFunction(upFunction).
	    AddFunction(downFunction)

	return migrationGenerator.classBuilder.GetClass()
}