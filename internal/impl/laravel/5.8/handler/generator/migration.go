package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/helper"
	"github.com/iancoleman/strcase"
)

type MigrationGenerator struct {
	api.Generator
	classBuilder interfaces.Class
	tableName    string
	columns      []*core.SimpleStatement
	softDeletes  bool
	auditCols    bool
	pkColVal     string
	timestamps   bool
}

const defaultColVal		= "unsignedBigInteger"
const createdByStr 		= "created_by"
const updatedByStr 		= "updated_by"
const timestampCol 		= `$table->timestamps()`
const softDeletesCol 	= `$table->softDeletes()`

/**
Creates a new instance of this generator with a new interfaces.Class
*/
func NewMigrationGenerator() *MigrationGenerator {
	return &MigrationGenerator{
		classBuilder: builder.NewClassBuilder(),
		columns:      []*core.SimpleStatement{},
		softDeletes:  false,
		auditCols:    false,
		timestamps:   false,
		pkColVal:     defaultColVal,
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
	migrationGenerator.tableName = tableName
	return migrationGenerator
}

/**
 Adds the given column to the up function of this migration.
 Parameters:
	- column: core.SimpleStatement of the column to add.
 Returns:
	- instance of the migration generator object.
 Example:
	- AddColumn(core.NewSimpleStatement('$this->string('name')->unique()'))
*/
func (migrationGenerator *MigrationGenerator) AddColumn(column core.SimpleStatement) *MigrationGenerator {
	return migrationGenerator.AddColumns([]*core.SimpleStatement{&column})
}

/**
 Adds the given list of columns to the up function of this migration.
 Parameters:
	- columns: A list of core.SimpleStatement of the columns to add.
 Returns:
	- instance of the migration generator object.
 Example:
	- AddColumns([]core.SimpleStatement{
		core.NewSimpleStatement('$this->string('name')->unique()')
		core.NewSimpleStatement('$this->string('phone_number', 12)->unique()')
	  })
*/
func (migrationGenerator *MigrationGenerator) AddColumns(columns []*core.SimpleStatement) *MigrationGenerator {
	migrationGenerator.columns = append(migrationGenerator.columns, columns...)
	return migrationGenerator
}

/**
 Sets the AuditCols field of this generator. During build adds string such as
 `$table->unsignedBigInteger('created_by');` and `$table->unsignedBigInteger('updated_by');`
 to the migration. The method name changes depending upon the col type used in the users table.
 NOTE- If this is set and SetPkCol is not used `unsignedBigInteger` is used as the default.
 Parameters
	- auditCols:	bool 		If set this generator adds created by and updated by cols to the migration file
 Returns:
	- instance of the generator object
 Example:
	- builder.SetAuditCols(true)
*/
func (migrationGenerator *MigrationGenerator) SetAuditCols(auditCols bool) *MigrationGenerator {
	migrationGenerator.auditCols = auditCols
	return migrationGenerator
}

/**
 Sets the pkCol field of this generator. During build adds string such as
 `$table->unsignedBigInteger('created_by');` and `$table->unsignedBigInteger('updated_by');`
 NOTE if this is not set the default value (`unsignedBigInteger`) is used.
 Parameters
	- pkColType:	string		The primary key col type of users table
 Returns:
	- instance of the generator object
 Example:
	- builder.SetPkCol(true)
*/
func (migrationGenerator *MigrationGenerator) SetPkCol(pkCol string) *MigrationGenerator {
	migrationGenerator.pkColVal = pkCol
	return migrationGenerator
}

/**
 Sets the timestamps field of this generator. During build adds the string `$table->softDeletes();`
 to the migration.
 Returns:
	- instance of the generator object
 Example:
	- builder.SetSoftDeletes(true)
*/
func (migrationGenerator *MigrationGenerator) SetSoftDeletes(softDeletes bool) *MigrationGenerator {
	migrationGenerator.softDeletes = softDeletes
	return migrationGenerator
}

/**
 Sets the timestamps field of this generator. During build adds the string `$table->timestamps();`
 to the migration.
 Returns:
	- instance of the generator object
 Example:
	- builder.SetTimestamps(true)
*/
func (migrationGenerator *MigrationGenerator) SetTimestamps(timestamps bool) *MigrationGenerator {
	migrationGenerator.timestamps = timestamps
	return migrationGenerator
}

/**
 Generates corresponding core.Class from the given parameters of this migration.
 Returns:
	- a pointer to the corresponding core.Class of this migration
*/
func (migrationGenerator *MigrationGenerator) Build() *core.Class {
	// Preparing the arguments for up function
	if migrationGenerator.softDeletes {
		migrationGenerator.AddColumn(*core.NewSimpleStatement(softDeletesCol))
	}

	if migrationGenerator.auditCols {
		cbstr := helper.ColTypeSwitcher(migrationGenerator.pkColVal, createdByStr, []string{})
		upstr := helper.ColTypeSwitcher(migrationGenerator.pkColVal, updatedByStr, []string{}) + "->nullable()"
		migrationGenerator.AddColumns([]*core.SimpleStatement{
			core.NewSimpleStatement(cbstr), core.NewSimpleStatement(upstr),
		})
	}

	if migrationGenerator.timestamps {
		migrationGenerator.AddColumns([]*core.SimpleStatement{
			core.NewSimpleStatement(timestampCol),
		})
	}

	arg1 := api.TabbedUnit(core.NewParameter("'" + migrationGenerator.tableName + "'"))
	closure := builder.NewFunctionBuilder().AddArgument("Blueprint $table")
	for _, element := range migrationGenerator.columns {
		closure.AddStatement(element)
	}

	arg2 := api.TabbedUnit(closure.GetFunction())

	// Preparing the statements for up function
	schemaBlock := api.TabbedUnit(core.NewFunctionCall("Schema::create").AddArg(&arg1).AddArg(&arg2))
	upFunction := builder.NewFunctionBuilder().SetName("up").SetVisibility("public").
		AddStatement(schemaBlock).GetFunction()

	// Preparing the statements for down function
	dropStatement := api.TabbedUnit(core.NewSimpleStatement("Schema::dropIfExists('" + migrationGenerator.tableName + "')"))
	downFunction := builder.NewFunctionBuilder().SetName("down").SetVisibility("public").
		AddStatement(dropStatement).GetFunction()

	migrationGenerator.classBuilder.AddImports([]string{
		`Illuminate\Database\Migrations\Migration`,
		`Illuminate\DatabaseSchema\Blueprint`,
		`Illuminate\Support\Facades\Schema`,
	}).SetExtends("Migration").AddFunction(upFunction).
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
