package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"fmt"
	"strings"
)

type RouteGenerator struct {
	api.Generator
	imports []api.TabbedUnit
	routes  []*core.FunctionCall
}

func NewRouteGenerator() *RouteGenerator {
	return &RouteGenerator{
		imports: []api.TabbedUnit{},
		routes:  []*core.FunctionCall{},
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
	- identifier: name of the model for which routes are to be generated
	-controller : configuration for controller
Returns:
	- instance of the generator object
Sample Usage:
	Only that routes will be added that are in the supported methods array of controller config
	- AddDefaultRestRoutes('Order')
	routes generated for the above call are
	Route::get(/order/{id}, OrderController@findById);
	Route::get(/order/all, OrderController@getAll);
	Route::post(/order/create, OrderController@create);
	Route::post(/order/edit/{id}, OrderController@edit);
	Route::post(/order/delete/{id}, OrderController@delete);
*/
func (routeGenerator *RouteGenerator) AddDefaultRestRoutes(modelName string, supportedHttpMethods []string) *RouteGenerator {

	type RouteConfig struct {
		method         string
		actionFunction string
		subURI         string
		httpMethod     string
	}

	var apiRouteConfig = []RouteConfig{
		{actionFunction: MethodNameCreate, method: HttpPost, subURI: ""},
		{actionFunction: MethodNameUpdate, method: HttpPut, subURI: "{id}"},
		{actionFunction: MethodNameDelete, method: HttpDelete, subURI: "{id}"},
		{actionFunction: MethodNameFindById, method: HttpGet, subURI: "{id}"},
		{actionFunction: MethodNameGetAll, method: HttpGet, subURI: "all"},
	}

	for _, routeConfig := range apiRouteConfig {

		/*CHECK IF METHOD IS PRESENT IN SUPPORTEDMETHODS ARRAY OF CONTROLLER IN JSON FILE
		IF PRESENT ADD THE METHOD ELSE DONT ADD*/
		if contains(supportedHttpMethods, routeConfig.method) || len(supportedHttpMethods) == 0 {
			uri := "/" + strings.ToLower(modelName)
			if routeConfig.subURI != "" {
				uri = uri + "/" + routeConfig.subURI
			}
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
	uri = `"` + uri + `"`
	action = `"` + action + `"`
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
	buildRoutFile := []api.TabbedUnit{core.NewSimpleStatement(`use Illuminate\Support\Facades\Route`)}

	/*ADD ALL FUNCTION CALLS*/
	for _, functionCall := range routeGenerator.routes {
		buildRoutFile = append(buildRoutFile, functionCall)
	}

	return buildRoutFile
}

/**
Returns:
	- contents of route file in string file
Sample Usage:
	-eg.output:
	use Illuminate\Support\Facades\Route;
	Route::get(/order/{id}, OrderController@fetchById);
	Route::get(/order/all, OrderController@getAll);
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.ToLower(a) == strings.ToLower(e) {
			return true
		}
	}
	return false
}
