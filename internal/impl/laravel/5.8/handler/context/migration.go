package context

import (
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/generator"
)

type Migration struct {
	BaseContext
	migrationContext map[string]*MigrationInfo
}

type MigrationInfo struct {
	Class         *generator.MigrationGenerator
	PrimaryKeyCol []string
}

func (info *MigrationInfo) AppendToPk(col string) {
	info.PrimaryKeyCol = append(info.PrimaryKeyCol, col)
}

func NewMigrationContext() *Migration {
	return &Migration{migrationContext: make(map[string]*MigrationInfo)}
}

/**
Store a MigrationInfo instance.
*/
func (migration *Migration) AddToCtx(key string, value interface{}) {
	migration.migrationContext[key] = &MigrationInfo{Class: value.(*core.Class), PrimaryKeyCol: []string{}}
}

/**
Fetches a MigrationInfo instance
The user of this method must cast and fetch appropriate data
*/
func (migration *Migration) GetCtx(key string) interface{} {
	return migration.migrationContext[key]
}
