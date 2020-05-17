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
	return &RouteGenerator{}
}

func (routeGenerator *RouteGenerator) Routes() []*core.FunctionCall {
	return routeGenerator.routes
}

func (routeGenerator *RouteGenerator) AddResourceRoutes(modelName string) *RouteGenerator {

	type RouteConfig struct {
		method         string
		actionFunction string
		subURI         string
	}

	var apiRouteConfig = []RouteConfig{
		{actionFunction: "get-by-id", method: "get", subURI: "get-by-id/{id}"},
		{actionFunction: "get-all", method: "get", subURI: "get-all"},
		{actionFunction: "create", method: "post", subURI: "create"},
		{actionFunction: "edit", method: "post", subURI: "edit/{id}"},
		{actionFunction: "delete", method: "post", subURI: "delete/{id}"},
	}

	for _, routeConfig := range apiRouteConfig {
		uri := "/" + strings.ToLower(modelName) + "/" + routeConfig.subURI
		action := modelName + "Controller" + "@" + routeConfig.actionFunction
		routeGenerator.AddRoute(routeConfig.method, uri, action)
	}

	return routeGenerator
}

func (routeGenerator *RouteGenerator) AddRoute(method string, uri string, action string) *RouteGenerator {
	route := core.NewFunctionCall("Route::" + method)
	uriTabbedUnit := api.TabbedUnit(core.NewParameter(uri))
	actionTabbedUnit := api.TabbedUnit(core.NewParameter(action))
	route.AddArg(&uriTabbedUnit).AddArg(&actionTabbedUnit)
	routeGenerator.routes = append(routeGenerator.Routes(), route)
	return routeGenerator
}

func (routeGenerator *RouteGenerator) Build() []*api.TabbedUnit {
	var buildedRouteFile []*api.TabbedUnit

	importStmt := api.TabbedUnit(core.NewSimpleStatement(`use Illuminate\Support\Facades\Route`))
	buildedRouteFile = append(buildedRouteFile, &importStmt)

	/*ADD ALL FUNCTION CALLS*/
	for _, functionCall := range routeGenerator.routes {
		funCall := api.TabbedUnit(functionCall)
		buildedRouteFile = append(buildedRouteFile, &funCall)
	}

	return buildedRouteFile
}

func (routeGenerator *RouteGenerator) String() string {
	var builder strings.Builder

	builderRouteFile := routeGenerator.Build()

	for _, element := range builderRouteFile {
		fmt.Fprint(&builder, *element, "\n")
	}

	return builder.String()
}
