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
		{genModel("Hello"), api.FromContext(context.ContextModel,
			"Hello").(*generator.ModelGenerator)},
		{genModel("World"), api.FromContext(context.ContextModel,
			"World").(*generator.ModelGenerator)},
	}
	for _, element := range classes {
		if element.modelMigExpected != element.modelMigOut {
			t.Error("Unexpected data")
		}
	}
	if nil != api.FromContext(context.ContextModel, "nonexistentRecords") {
		t.Error("Unexpected data")
	}

}

func genModel(className string) *generator.ModelGenerator {
	modelGen := generator.NewModelGenerator().SetName(className)
	context.GetFromRegistry("model").AddToCtx(className, modelGen)
	return modelGen
}

