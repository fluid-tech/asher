package handler

import (
	"asher/internal/impl/laravel/5.8/handler"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/models"
	"testing"
)


func TestController(t *testing.T) {


	RESTControllerConfigWithNoHttpMethods := models.Controller{
		Rest:        true,
		Mvc:         false,
		HttpMethods: []string{"GET"},
		Type:        "",
	}

	RESTControllerConfigWithAllHttpMethodsAndFileType := models.Controller{
		Rest:        true,
		Mvc:         false,
		HttpMethods: []string{"GET","POST","PATCH","DELETE"},
		Type:        "file",
	}

	RESTControllerConfigWithGETMethodAndImageType := models.Controller{
		Rest:        true,
		Mvc:         false,
		HttpMethods: []string{"GET"},
		Type:        "image",
	}

	var table = []*struct {
		in  []string
		out []string
	}{
		{genControllerTest("Order1", RESTControllerConfigWithNoHttpMethods,t,true), []string{CTOrder1Controller,CTOrder1Transactor,CTOrder1Mutator,CTOrder1Query,CTRouteFileAfterOrder1}},

		{genControllerTest("Order2", RESTControllerConfigWithAllHttpMethodsAndFileType, t,false), []string{CTOrder2Controller,CTOrder2Transactor,CTOrder2Mutator,CTOrder2Query,CTRouteFileAfterOrder2}},

		{genControllerTest("Order3", RESTControllerConfigWithGETMethodAndImageType,t,false), []string{CTOrder3Controller,CTOrder3Transactor,CTOrder3Mutator,CTOrder3Query,CTRouteFileAfterOrder3}},
	}

	//if table[0].in[0] == ""  {
	//	t.Error("controller handler returned Nothing")
	//}

	for _, element := range table {
		for j:=0 ; j<5 ;j++ {
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
		t.Error(" first call controller handler didi not returned route file")
	}

	/*Second and greater controllerHandlerCall should return only 3 files*/
	if !isFirstCall && !(len(emitterFiles) == 4) {
		t.Error("Not returend 3 files returned ",len(emitterFiles))
	}

	retrivedControllerGen := fromControllerReg(className)
	retrivedTransactorGen :=fromTransactorReg(className)
	retrivedMutatorGen :=fromMutattorReg(className)
	retrivedRouteGen :=fromRouteReg("api")
	retrivedQueryGen :=fromQueryReg(className)

	//fmt.Print(retrivedControllerGen)
	//fmt.Print(retrivedTransactorGen)
	//fmt.Print(retrivedMutatorGen)
	//fmt.Print(retrivedRouteGen)
	//fmt.Print(retrivedQueryGen)
	//
	//fmt.Print("--------------------------------------------------------------")


	if retrivedControllerGen == nil {
		t.Errorf("controller for %s doesnt exist ", className)
	}

	if retrivedTransactorGen == nil {
		t.Errorf("transactor for %s doesnt exist ", className)
	}

	if retrivedMutatorGen == nil {
		t.Errorf("mutator for %s doesnt exist ", className)
	}

	if retrivedRouteGen == nil {
		t.Errorf("route for %s doesnt exist ", className)
	}

	if retrivedQueryGen == nil {
		t.Errorf("query for %s doesnt exist ", className)
	}

	return []string{retrivedControllerGen.String(), retrivedTransactorGen.String(), retrivedMutatorGen.String(), retrivedQueryGen.String(),retrivedRouteGen.String()}
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



