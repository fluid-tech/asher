package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"fmt"
	"testing"
)

//Route::get('/user', 'UserController@index');
func TestRouteGenerator(t *testing.T) {
	apiGenerator:=generator.NewRouteGenerator("api")
	apiGenerator.AddResourceRoutes("Order")
	fmt.Print(apiGenerator)
}

func testGenResourceRoute(t *testing.T){

}

func testGenCustomRoute(t *testing.T){

}