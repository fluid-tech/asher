package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestModelGenerator(t *testing.T) {
	var emptyArr []string
	emptyMap := map[string]string{}
	fillableArr := []string{"name", "phone_number"}
	hiddenFields := []string{"password", "gender"}
	createRules := map[string]string{

		"name":         "string|max:255|unique",
		"phone_number": "string|max:12|unique:users,id",
	}
	updateRules := map[string]string{
		"name":         "string|max:255|unique",
		"phone_number": "string|max:12|unique:users,id",
	}
	createRules2 := map[string]string{
		"name":         "string|max:255|unique",
		"phone_number": "string|max:12|unique:users",
	}
	updateRules2 := map[string]string{
		"name":         "string|max:255|unique",
		"phone_number": "string|max:12|unique:users",
	}
	var table = []*api.GeneralTest{
		genModelGeneratorTest("student_allotments", emptyArr, emptyArr, emptyMap, emptyMap, EmptyModel),
		genModelGeneratorTest("student_allotments", fillableArr, emptyArr, emptyMap, emptyMap, ModelWithFillable),
		genModelGeneratorTest("student_allotments", emptyArr, hiddenFields, emptyMap, emptyMap, ModelWithHidden),
		genModelGeneratorTest("student_allotments", emptyArr, emptyArr, createRules, emptyMap,
			ModelWithCreateValidationRules),
		genModelGeneratorTest("student_allotments", emptyArr, emptyArr, emptyMap, updateRules,
			ModelWithUpdateValidationRules),
		genModelGeneratorTest("student_allotments", emptyArr, emptyArr, createRules2, updateRules2,
			ModelWithUpdateValidationRulesWithoutId),
	}

	api.IterateAndTest(table, t)
}

/**
A helper function to generate GeneralTest cases for ModelGenerator
*/
func genModelGeneratorTest(name string, fillables []string, hiddenFields []string, createRules map[string]string,
	updateRules map[string]string, expectedCode string) *api.GeneralTest {
	modelGenerator := generator.NewModelGenerator().SetName(name)
	for _, fillable := range fillables {
		modelGenerator.AddFillable(fillable)
	}
	for _, hiddenField := range hiddenFields {
		modelGenerator.AddHiddenField(hiddenField)
	}
	for column, rules := range createRules {
		modelGenerator.AddCreateValidationRule(column, rules, name)
	}
	for column, rules := range updateRules {
		modelGenerator.AddUpdateValidationRule(column, rules, name)
	}
	stringer := modelGenerator.Build().String()
	return api.NewGeneralTest(stringer, expectedCode)
}
