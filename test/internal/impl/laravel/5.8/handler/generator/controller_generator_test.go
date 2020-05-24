package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestBasicController(t *testing.T) {
	restWithCreate := []string{"post"}
	restWithUpdate := []string{"put"}
	restWithGet := []string{"get"}
	restWithDelete := []string{"DELETE"}
	restWithAllFUnctions := []string{"POST", "PUT", "DELETE", "GET"}
	table := []*api.GeneralTest{
		genControllerGeneratorTest(nil, AllFunctionsRestController, "Order"),
		genControllerGeneratorTest(restWithCreate, CreateRestController,"Order"),
		genControllerGeneratorTest(restWithUpdate, UpdateRestController,"Order"),
		genControllerGeneratorTest(restWithDelete, DeleteFunctionRestController,"Order"),
		genControllerGeneratorTest(restWithGet, GetFUnctionRestController,"Order"),
		genControllerGeneratorTest(restWithAllFUnctions, AllFunctionsRestController,"Order"),
		genControllerGeneratorTest(restWithAllFUnctions,StudentController,"Student"),
		genControllerGeneratorTest(restWithGet,TeacherController,"Teacher"),
		genControllerGeneratorTest([]string{"POST","DELETE","PUT"},AdminController,"Admin"),
	}

	api.IterateAndTest(table, t)

}

func genControllerGeneratorTest(array []string, expectedCodeString string, modelName string) *api.GeneralTest {
	conGen := generator.NewControllerGenerator()
	conGen.SetIdentifier(modelName)
	conGen.AddFunctionsInController(array)
	return api.NewGeneralTest(conGen.String(), expectedCodeString)
}
