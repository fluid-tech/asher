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
	- instance of the migration generator object
*/
func (migrationGenerator *MigrationGenerator) SetName(columnName string) *MigrationGenerator {
	className := "Create" + strcase.ToCamel(columnName) + "Table"
	migrationGenerator.classBuilder.SetName(className)
	return migrationGenerator
}

func (migrationGenerator *MigrationGenerator) AddColumn() {}


func (migrationGenerator *MigrationGenerator) Build() *core.Class {

}