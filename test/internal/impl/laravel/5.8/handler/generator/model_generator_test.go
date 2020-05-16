package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestModelGenerator(t *testing.T) {
	var table = []*api.GeneralTest{
		getEmptyModel(),
		getModelWithFillable(),
		getModelWithHiddenFields(),
		getModelWithCreateValidationRules(),
		getModelWithUpdateValidationRules(),
	}

	api.IterateAndTest(table, t)
}

func getEmptyModel() *api.GeneralTest {
	modelGenerator := generator.NewModelGenerator().SetName("student_allotments")
	return api.NewGeneralTest(modelGenerator.Build().String(), EmptyModel)
}

func getModelWithFillable() *api.GeneralTest {
	modelGenerator := generator.NewModelGenerator().SetName("student_allotments").
		AddFillable("name").AddFillable("phone_number")
	return api.NewGeneralTest(modelGenerator.Build().String(), ModelWithFillable)
}

func getModelWithHiddenFields() *api.GeneralTest {
	modelGenerator := generator.NewModelGenerator().SetName("student_allotments").
		AddHiddenField("password").AddHiddenField("gender")
	return api.NewGeneralTest(modelGenerator.Build().String(), ModelWithHidden)
}

func getModelWithCreateValidationRules() *api.GeneralTest {
	modelGenerator := generator.NewModelGenerator().SetName("student_allotments").
		AddCreateValidationRule("name", "string|max:255|unique:users").
		AddCreateValidationRule("phone_number", "string|max:12|unique:users")
	return api.NewGeneralTest(modelGenerator.Build().String(), ModelWithCreateValidationRules)
}

func getModelWithUpdateValidationRules() *api.GeneralTest {
	modelGenerator := generator.NewModelGenerator().SetName("student_allotments").
		AddUpdateValidationRule("name", "string|max:255|unique:users").
		AddUpdateValidationRule("phone_number", "string|max:12|unique:users")
	return api.NewGeneralTest(modelGenerator.Build().String(), ModelWithUpdateValidationRules)
}
