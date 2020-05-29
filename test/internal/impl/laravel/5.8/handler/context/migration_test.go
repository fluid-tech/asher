package context

import (
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestMigrationContext(t *testing.T) {
	var classes = []struct {
		MigrationMigOut      *generator.MigrationGenerator
		MigrationMigExpected *generator.MigrationGenerator
	}{

		{genMigration("Hello"), api.FromContext(context.Migration,
			"Hello").(*generator.MigrationGenerator)},

		{genMigration("World"), api.FromContext(context.Migration,
			"World").(*generator.MigrationGenerator)},
	}
	for _, element := range classes {
		if element.MigrationMigExpected != element.MigrationMigOut {
			t.Error("Unexpected data")
		}
	}
	if nil != api.FromContext(context.Migration, "nonexistentRecords") {
		t.Error("Unexpected data")
	}
}

func genMigration(className string) *generator.MigrationGenerator {
	MigrationGen := generator.NewMigrationGenerator().SetName(className)
	context.GetFromRegistry(context.Migration).AddToCtx(className, MigrationGen)
	return MigrationGen
}
