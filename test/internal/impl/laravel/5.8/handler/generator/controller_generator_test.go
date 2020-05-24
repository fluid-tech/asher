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
		genControllerGeneratorTest(nil, AllFunctionsRestController),
		genControllerGeneratorTest(restWithCreate, CreateRestController),
		genControllerGeneratorTest(restWithUpdate, UpdateRestController),
		genControllerGeneratorTest(restWithDelete, DeleteFunctionRestController),
		genControllerGeneratorTest(restWithGet, GetFUnctionRestController),
		genControllerGeneratorTest(restWithAllFUnctions, AllFunctionsRestController),
	}

	api.IterateAndTest(table, t)


}

func genControllerGeneratorTest(array []string, expectedCodeString string) *api.GeneralTest {
	conGen := generator.NewControllerGenerator()
	conGen.SetIdentifier("Order")
	conGen.AddFunctionsInController(array)
	//fmt.Print(conGen)
	return api.NewGeneralTest(conGen.String(), expectedCodeString)
}
