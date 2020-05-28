package generator

import "asher/internal/api/codebuilder/php/core"

type TransactorMigration struct {
	migrationGen *MigrationGenerator
}

func NewTransactorMigration(migrationGen *MigrationGenerator) *TransactorMigration {
	return &TransactorMigration{migrationGen: migrationGen}
}

func (transactorMigrationGen *TransactorMigration) AddMigrationForFileUrls() *TransactorMigration {
	transactorMigrationGen.migrationGen.AddColumn(core.NewSimpleStatement(`$table->longText('file_urls')->nullable()`))
	return transactorMigrationGen
}
