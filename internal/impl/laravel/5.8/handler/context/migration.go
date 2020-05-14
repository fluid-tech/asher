package context

import (
	"asher/internal/api/codebuilder/php/core"
)

type Migration struct {
	BaseContext
	migrationContext map[string]*MigrationInfo
}

type MigrationInfo struct {
	Class         *core.Class
	PrimaryKeyCol string
}

func NewMigrationContext() *Migration {
	return &Migration{migrationContext: make(map[string]*MigrationInfo)}
}

/**
Store a MigrationInfo instance.
 */
func (migration *Migration) AddToCtx(key string, value interface{}) {
	migration.migrationContext[key] = &MigrationInfo{Class: value.(*core.Class), PrimaryKeyCol: ""}
}

/**
Fetches a MigrationInfo instance
The user of this method must cast and fetch appropriate data
 */
func (migration *Migration) GetCtx(key string) interface{} {
	return migration.migrationContext[key]
}
