package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/models"
	"fmt"
	"testing"
)

//Route::get('/user', 'UserController@index');
func TestRouteGenerator(t *testing.T) {
	apiGenerator := generator.NewRouteGenerator()

	apiGenerator.AddDefaultRestRoutes("Order", models.Controller{
		Rest:        false,
		Mvc:         false,
		HttpMethods: []string{"POST"},
		Type:        "",
	})
	///*ADDS SET OF DEFAULT API ROUTES*/
	//apiGenerator.AddDefaultRestRoutes("Order")
	//
	///*ADDS A SPECIFIC ROUTE*/
	//apiGenerator.AddRoute("get", `"/order-products"`, "OrderController@getAll")
	//apiGenerator.AddRoute("get", "/order-products", "OrderController@getAll")
	//apiGenerator.AddRoute("get", "/order-products", "OrderController@getAll")

	fmt.Print(apiGenerator)
}

func testGenResourceRoute(t *testing.T) {

}

func testGenCustomRoute(t *testing.T) {

}
