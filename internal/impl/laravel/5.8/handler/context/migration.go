package context

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
)

type Migration struct {
	BaseContext
	migrationGenerator map[string]*generator.MigrationGenerator
}

func NewMigrationContext() *Migration {
	return &Migration{migrationGenerator: make(map[string]*generator.MigrationGenerator)}
}

/**
Store a MigrationInfo instance.
*/
func (migration *Migration) AddToCtx(key string, value interface{}) {
	migration.migrationGenerator[key] = value.(*generator.MigrationGenerator)
}

/**
Fetches a MigrationInfo instance
The user of this method must cast and fetch appropriate data
*/
func (migration *Migration) GetCtx(key string) interface{} {
	return migration.migrationGenerator[key]
}
