package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/models"
	"fmt"
	"strings"
)

type RouteGenerator struct {
	imports []api.TabbedUnit
	routes  []*core.FunctionCall
}

func NewRouteGenerator() *RouteGenerator {
	return &RouteGenerator{
		imports:	[]api.TabbedUnit{},
		routes: 	[]*core.FunctionCall{},
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

/**
Returns the array of functional calls for every model to add their routes
Returns:
	- array of *core.FunctionCall
*/
func (routeGenerator *RouteGenerator) Routes() []*core.FunctionCall {
	return routeGenerator.routes
}

/**
Add the predefined sets of rest routes for the specific model in the generator routes array
Parameters:
	- modelName: name of the model for which routes are to be generated
	-controller : configuration for controller
Returns:
	- instance of the generator object
Sample Usage:
	Only that routes will be added thet are in the supported methods array of controller config
	- AddDefaultRestRoutes('Order')
	routes generated for the above call are
	Route::get(/order/get-by-id/{id}, OrderController@get-by-id);
	Route::get(/order/get-all, OrderController@get-all);
	Route::post(/order/create, OrderController@create);
	Route::post(/order/edit/{id}, OrderController@edit);
	Route::post(/order/delete/{id}, OrderController@delete);
*/
func (routeGenerator *RouteGenerator) AddDefaultRestRoutes(modelName string, controller models.Controller) *RouteGenerator {

	type RouteConfig struct {
		method         string
		actionFunction string
		subURI         string
		httpMethod 	string
	}

	var apiRouteConfig = []RouteConfig{
		{actionFunction: "get-by-id", method: "GET", subURI: "{id}"},
		{actionFunction: "all", method: "GET", subURI: "all"},
		{actionFunction: "create", method: "POST", subURI: "create"},
		{actionFunction: "edit", method: "PATCH", subURI: "edit/{id}"},
		{actionFunction: "delete", method: "DELETE", subURI: "delete/{id}"},
	}

	for _, routeConfig := range apiRouteConfig {

		/*CHECK IF METHOD IS PRESENT IN SUPPORTEDMETHODS ARRAY OF CONTROLLER IN JSON FILE
		IF PRESENT ADD THE METHOD ELSE DONT ADD*/
		if contains(controller.supportedMethods, routeConfig.method) {
			uri := "/" + strings.ToLower(modelName) + "/" + routeConfig.subURI
			action := modelName + "Controller" + "@" + routeConfig.actionFunction
			routeGenerator.AddRoute(strings.ToLower(routeConfig.method), uri, action)
		}

	}

	return routeGenerator
}

/**
Adds the specific routes to the routes array of generator
Parameters:
	- method: http methods like get,post,delete...
	- uri: uri to perform the action
	-action: ControllerName@functionName
Returns:
	- instance of the generator object
Sample Usage:
	- AddRoute("post","/order/create","OrderController@create")
	Route::post(/order/create, OrderController@create);
*/
func (routeGenerator *RouteGenerator) AddRoute(method string, uri string, action string) *RouteGenerator {
	route := core.NewFunctionCall("Route::" + method)
	route.AddArg(core.NewParameter(uri)).AddArg(core.NewParameter(action))
	routeGenerator.routes = append(routeGenerator.Routes(), route)
	return routeGenerator
}

/**
Returns the array of tabbedUnits in which the imports array is followed by routes array
Returns:
	- array of tabbed units
*/
func (routeGenerator *RouteGenerator) Build() []api.TabbedUnit {
	var builtRouteFile []api.TabbedUnit

	importStmt := core.NewSimpleStatement(`use Illuminate\Support\Facades\Route`)
	builtRouteFile = append(builtRouteFile, importStmt)

	/*ADD ALL FUNCTION CALLS*/
	for _, functionCall := range routeGenerator.routes {
		builtRouteFile = append(builtRouteFile, functionCall)
	}

	return builtRouteFile
}

/**
Returns:
	- contents of route file in string file
Sample Usage:
	-eg.output:
	use Illuminate\Support\Facades\Route;
	Route::get(/order/get-by-id/{id}, OrderController@get-by-id);
	Route::get(/order/get-all, OrderController@get-all);
	Route::post(/order/create, OrderController@create);
	Route::post(/order/edit/{id}, OrderController@edit);
	Route::post(/order/delete/{id}, OrderController@delete);
*/
func (routeGenerator *RouteGenerator) String() string {
	var builder strings.Builder

	builderRouteFile := routeGenerator.Build()

	for _, element := range builderRouteFile {
		fmt.Fprint(&builder, element, "\n")
	}

	return builder.String()
}
