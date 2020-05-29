package generator

import "asher/internal/api/codebuilder/php/core"

const fileUrlsMigration = `$table->longText('file_urls')->nullable()`

/*Migration related activities for transactor will be handled in this class*/
type TransactorMigration struct {
	migrationGen *MigrationGenerator
}

/**
Returns a New Instance of TransactorModel
Parameters:
	- migrationGen: migration generator on which transactor specific activities will be carried out
Returns:
	- instance of the generator object
*/
func NewTransactorMigration(migrationGen *MigrationGenerator) *TransactorMigration {
	return &TransactorMigration{migrationGen: migrationGen}
}

/**
To store file urls a column should be added through migration, used for image and file transactor
Returns:
	- instance of the generator object
*/
func (transactorMigrationGen *TransactorMigration) AddMigrationForFileUrls() *TransactorMigration {
	transactorMigrationGen.migrationGen.AddColumn(core.NewSimpleStatement(fileUrlsMigration))
	return transactorMigrationGen
}


