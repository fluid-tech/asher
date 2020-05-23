package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"strings"
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
	if array != nil {
		if len(array) >= 0 {
			conGen.AddConstructorFunction()
			for _, element := range array {
				switch strings.ToLower(element) {
				case "post":
					conGen.AddCreateFunction()
				case "get":
					conGen.AddFindByIdFunction().AddGetAllFunction()
				case "put":
					conGen.AddUpdateFunction()
				case "delete":
					conGen.AddDeleteFunction()
				}
			}
		}
	}
	return api.NewGeneralTest(conGen.String(), expectedCodeString)
}
