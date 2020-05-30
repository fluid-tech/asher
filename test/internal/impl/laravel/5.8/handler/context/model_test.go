package context

import (
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestModelContext(t *testing.T) {
	var classes = []struct {
		modelMigOut      *generator.ModelGenerator
		modelMigExpected *generator.ModelGenerator
	}{
		{genModel("Hello"), api.FromContext(context.Model,
			"Hello").(*generator.ModelGenerator)},
		{genModel("World"), api.FromContext(context.Model,
			"World").(*generator.ModelGenerator)},
	}
	for _, element := range classes {
		if element.modelMigExpected != element.modelMigOut {
			t.Error("Unexpected data")
		}
	}
	if nil != api.FromContext(context.Model, "nonexistentRecords") {
		t.Error("Unexpected data")
	}

}

func genModel(className string) *generator.ModelGenerator {
	modelGen := generator.NewModelGenerator().SetName(className)
	context.GetFromRegistry(context.Model).AddToCtx(className, modelGen)
	return modelGen
}
