package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"fmt"
	"testing"
)

func TestController(t *testing.T) {
	controllerGenerator := generator.NewControllerGenerator()
	controllerGenerator.SetIdentifier("Order")
	fmt.Print(controllerGenerator.String())
	//test := api.NewGeneralTest(controllerGenerator.String(), BasicTransactor)
	//api.IterateAndTest([]*api.GeneralTest{
	//	test,
	//},t)
}
