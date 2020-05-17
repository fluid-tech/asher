package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"fmt"
	"testing"
)

func TestRouteGenerator(t *testing.T) {
	apiGenerator:=generator.NewRouteGenerator("api")
	apiGenerator.AddResourceRoutes("Order")
	emmitterFile:=apiGenerator.Build()
	if len(apiGenerator.Routes()) == 0 {
		t.Error("Values not inserted")
	}
	fmt.Print(emmitterFile.String())
}

func testGenResourceRoute(t *testing.T){

}

func testGenCustomRoute(t *testing.T){

}