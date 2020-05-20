package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"fmt"
	"strings"
)

type RouteGenerator struct {
	imports []*api.TabbedUnit
	routes  []*core.FunctionCall
}

func NewRouteGenerator() *RouteGenerator {
	return &RouteGenerator{
		imports:	[]*api.TabbedUnit{},
		routes: 	[]*core.FunctionCall{},
	}
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
Returns:
	- instance of the generator object
Sample Usage:
	- AddDefaultRestRoutes('Order')
	routes generated for the above call are
	Route::get(/order/get-by-id/{id}, OrderController@get-by-id);
	Route::get(/order/get-all, OrderController@get-all);
	Route::post(/order/create, OrderController@create);
	Route::post(/order/edit/{id}, OrderController@edit);
	Route::post(/order/delete/{id}, OrderController@delete);
*/
func (routeGenerator *RouteGenerator) AddDefaultRestRoutes(modelName string) *RouteGenerator {

	type RouteConfig struct {
		method         string
		actionFunction string
		subURI         string
	}

	var apiRouteConfig = []RouteConfig{
		{actionFunction: "get-by-id", method: "get", subURI: "{id}"},
		{actionFunction: "all", method: "get", subURI: "all"},
		{actionFunction: "create", method: "post", subURI: "create"},
		{actionFunction: "edit", method: "patch", subURI: "edit/{id}"},
		{actionFunction: "delete", method: "delete", subURI: "delete/{id}"},

	}

	for _, routeConfig := range apiRouteConfig {
		uri := "/" + strings.ToLower(modelName) + "/" + routeConfig.subURI
		action := modelName + "Controller" + "@" + routeConfig.actionFunction
		routeGenerator.AddRoute(routeConfig.method, uri, action)
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
	uriTabbedUnit := api.TabbedUnit(core.NewParameter(uri))
	actionTabbedUnit := api.TabbedUnit(core.NewParameter(action))
	route.AddArg(&uriTabbedUnit).AddArg(&actionTabbedUnit)
	routeGenerator.routes = append(routeGenerator.Routes(), route)
	return routeGenerator
}

/**
Returns the array of tabbedUnits in which the imports array is followed by routes array
Returns:
	- array of tabbed units
*/
func (routeGenerator *RouteGenerator) Build() []*api.TabbedUnit {
	var builtRouteFile []*api.TabbedUnit

	importStmt := api.TabbedUnit(core.NewSimpleStatement(`use Illuminate\Support\Facades\Route`))
	builtRouteFile = append(builtRouteFile, &importStmt)

	/*ADD ALL FUNCTION CALLS*/
	for _, functionCall := range routeGenerator.routes {
		funCall := api.TabbedUnit(functionCall)
		builtRouteFile = append(builtRouteFile, &funCall)
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
		fmt.Fprint(&builder, *element, "\n")
	}

	return builder.String()
}
