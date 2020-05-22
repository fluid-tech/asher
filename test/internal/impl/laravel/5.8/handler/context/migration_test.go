package context

import (
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"testing"
)

func TestMigrationContext(t *testing.T) {
	var classes = []struct {
		MigrationMigOut      *generator.MigrationGenerator
		MigrationMigExpected *generator.MigrationGenerator
	}{
		{genMigration("Hello"), fromMigrationReg("Hello")},
		{genMigration("World"), fromMigrationReg("World")},
		{ nil, fromMigrationReg("NonExistant")},
	}
	for _, element := range classes {
		if element.MigrationMigExpected != element.MigrationMigOut {
			t.Error("Unexpected data")
		}
	}
}

func genMigration(className string) *generator.MigrationGenerator {
	MigrationGen := generator.NewMigrationGenerator().SetName(className)
	context.GetFromRegistry("migration").AddToCtx(className, MigrationGen)
	return MigrationGen
}

func fromMigrationReg(className string) *generator.MigrationGenerator {
	return context.GetFromRegistry("migration").GetCtx(className).(*generator.MigrationGenerator)
}
