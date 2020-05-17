package context

import (
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"testing"
)

func TestMigrationContext(t *testing.T) {
	migration := context.GetFromRegistry("migration")
	var classes = []struct {
		class        *core.Class
		expectedName string
	}{
		{getClass("Hello"), "Hello"},
		{getClass("World"), "World"},
	}
	for _, element := range classes {
		migration.AddToCtx(element.class.Name, element.class)
		if migration.GetCtx(element.expectedName).(*generator.MigrationGenerator).Build().Name != element.expectedName {
			t.Error("Unexpected data")
		}
	}
}

func getClass(name string) *core.Class {
	class := core.Class{}
	class.Name = name
	class.Tabs = 0
	return &class
}
