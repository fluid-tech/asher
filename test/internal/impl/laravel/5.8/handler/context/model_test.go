package context

import (
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"testing"
)

func TestModelContext(t *testing.T) {
	var classes = []struct {
		modelMigOut      *generator.ModelGenerator
		modelMigExpected *generator.ModelGenerator
	}{
		{genModel("Hello"), fromModelReg("Hello")},
		{genModel("World"), fromModelReg("World")},
		{nil, fromModelReg("NonExistentRecord")},
	}
	for _, element := range classes {
		if element.modelMigExpected != element.modelMigOut {
			t.Error("Unexpected data")
		}
	}
}

func genModel(className string) *generator.ModelGenerator {
	modelGen := generator.NewModelGenerator().SetName(className)
	context.GetFromRegistry("model").AddToCtx(className, modelGen)
	return modelGen
}

func fromModelReg(className string) *generator.ModelGenerator {
	return context.GetFromRegistry("model").GetCtx(className).(*generator.ModelGenerator)
}
