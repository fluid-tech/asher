package handler

import (
	"asher/internal/impl/laravel/5.8/handler"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/models"
	generator2 "asher/test/internal/impl/laravel/5.8/handler/generator"
	"testing"
)

func TestController(t *testing.T) {

	/*Demo strings for Model Student have all HTTP methods  BASIC Transactor*/
	/*Demo strings for Model Teacher have all GET HTTP methods Teacher Image Transactor*/
	/*Demo strings for Model Admin have all PUT DELETE POST HTTP methods Admin FIle Transactor*/

	RESTControllerConfigWithALLHttpMethods := models.Controller{
		Rest:        true,
		Mvc:         false,
		HttpMethods: []string{},
		Type:        "default",
	}
	RESTControllerConfigWithGETHttpMethods := models.Controller{
		Rest:        true,
		Mvc:         false,
		HttpMethods: []string{"GET"},
		Type:        "image",
	}
	RESTControllerConfigWithPOSTPUTDELETEHttpMethods := models.Controller{
		Rest:        true,
		Mvc:         false,
		HttpMethods: []string{"POST", "DELETE", "PUT"},
		Type:        "file",
	}

	var table = []*struct {
		in  []string
		out []string
	}{
		{genControllerTest("Student", RESTControllerConfigWithALLHttpMethods, t, true),
			[]string{generator2.StudentController, generator2.StudentBasicTransactor, generator2.StudentBasicMutator, generator2.StudentBasicQuery,
				generator2.ApiRouteFileAfterStudentWithAllRoutes,generator2.StudentEmptyMigrationWithName, generator2.StudentEmptyModel}},

		{genControllerTest("Teacher", RESTControllerConfigWithGETHttpMethods, t, false),
			[]string{generator2.TeacherController, generator2.TeacherImageTransactor, generator2.TeacherBasicMutator, generator2.TeacherBasicQuery,
				generator2.ApiRouteFileAfterTeacherWithGetRoutes,generator2.TeacherMigrationForFileURLS,generator2.TeacherModelWithFileURLS}},

		{genControllerTest("Admin", RESTControllerConfigWithPOSTPUTDELETEHttpMethods, t, false),
			[]string{generator2.AdminController, generator2.AdminFileTransactor, generator2.AdminBasicMutator, generator2.AdminBasicQuery,
				generator2.ApiRouteFileAfterAdminWithPATCHPOSTDELTERoutes,generator2.AdminMigrationForFileURLS,generator2.AdminModelWithFileURLS}},
	}

	for _, element := range table {
		for j := 0; j < 7; j++ {
			if element.in[j] != element.out[j] {
				t.Errorf("in test case %d expected '%s' found '%s'", j, element.out[j], element.in[j])
			}
		}
	}

}

/**
 Returns a row indicating the following information
    - string of migration class
	- string of model class
*/
func genControllerTest(className string, controllerConfig models.Controller, t *testing.T, isFirstCall bool) []string {

	modelGen := generator.NewModelGenerator().SetName(className)
	migGen := generator.NewMigrationGenerator().SetName(className)

	context.GetFromRegistry("migration").AddToCtx(className, migGen)
	context.GetFromRegistry("model").AddToCtx(className, modelGen)

	emitterFiles, error := handler.NewControllerHandler().Handle(className, controllerConfig)

	if error != nil {
		t.Error(error)
	}

	if emitterFiles == nil {
		t.Error("controller handler returned Nothing")
	}

	/*only first call to handleController will return asher_api.php hence it will return 4 files*/
	if isFirstCall && !(len(emitterFiles) == 5) {
		t.Error(" first call controller handler did not returned route file")
	}

	/*Second and greater controllerHandlerCall should return only 3 files*/
	if !isFirstCall && !(len(emitterFiles) == 4) {
		t.Error("Not returned 4 files", len(emitterFiles))
	}

	retrievedControllerGen := fromControllerReg(className)
	retrievedTransactorGen := fromTransactorReg(className)
	retrievedMutatorGen := fromMutattorReg(className)
	retrievedRouteGen := fromRouteReg("api")
	retrievedQueryGen := fromQueryReg(className)

	if retrievedControllerGen == nil {
		t.Errorf("controller for %s doesnt exist ", className)
	}

	if retrievedTransactorGen == nil {
		t.Errorf("transactor for %s doesnt exist ", className)
	}

	if retrievedMutatorGen == nil {
		t.Errorf("mutator for %s doesnt exist ", className)
	}

	if retrievedRouteGen == nil {
		t.Errorf("route for %s doesnt exist ", className)
	}

	if retrievedQueryGen == nil {
		t.Errorf("query for %s doesnt exist ", className)
	}

	return []string{retrievedControllerGen.String(), retrievedTransactorGen.String(), retrievedMutatorGen.String(),
		retrievedQueryGen.String(), retrievedRouteGen.String(), migGen.String(), modelGen.String()}
}

func fromControllerReg(className string) *generator.ControllerGenerator {
	return context.GetFromRegistry("controller").GetCtx(className).(*generator.ControllerGenerator)
}

func fromTransactorReg(className string) *generator.TransactorGenerator {
	return context.GetFromRegistry("transactor").GetCtx(className).(*generator.TransactorGenerator)
}

func fromMutattorReg(className string) *generator.MutatorGenerator {
	return context.GetFromRegistry("mutator").GetCtx(className).(*generator.MutatorGenerator)
}

func fromQueryReg(className string) *generator.QueryGenerator {
	return context.GetFromRegistry("query").GetCtx(className).(*generator.QueryGenerator)
}

func fromRouteReg(className string) *generator.RouteGenerator {
	return context.GetFromRegistry("route").GetCtx(className).(*generator.RouteGenerator)
}
