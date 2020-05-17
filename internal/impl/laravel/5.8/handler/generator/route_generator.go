package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"fmt"
	"strings"
)

const apiRouteFileName = "asherapi.php"
const webRouteFileName = "asherweb.php"
const routeFilePath = "routes/"

type RouteGenerator struct {
	imports   []*api.TabbedUnit
	RouteFile *core.PhpEmitterFile
	routes    []*core.FunctionCall
}

func NewRouteGenerator(routeType string) *RouteGenerator {
	if routeType == "api" {
		temp:=make([]*api.TabbedUnit,1)
		return &RouteGenerator{
			RouteFile: core.NewPhpEmitterFile(apiRouteFileName, routeFilePath,temp, 1),
		}
	} else {
		return nil
	}
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
		{actionFunction: "create", method: "post", subURI: "create"},
		{actionFunction: "edit", method: "post", subURI: "edit/{id}"},
		{actionFunction: "delete", method: "post", subURI: "delete/{id}"},
		{actionFunction: "get-by-id", method: "get", subURI: "get-by-id/{id}"},
		{actionFunction: "get-all", method: "get", subURI: "get-all"},
	}

	for _, routeConfig := range apiRouteConfig {
		uri := "/" + strings.ToLower(modelName) + routeConfig.subURI
		action := modelName + "Controller" + "@" + routeConfig.actionFunction
		routeGenerator.AddRoute(routeConfig.method, uri, action)
	}

	return routeGenerator
}

func (routeGenerator *RouteGenerator) AddRoute(method string, uri string, action string) *RouteGenerator {
	route := core.NewFunctionCall("Route::" + method)
	uriTabbedUnit := api.TabbedUnit(core.NewSimpleStatement(uri))
	actionTabbedUnit := api.TabbedUnit(core.NewSimpleStatement(action))
	route.AddArg(&uriTabbedUnit).AddArg(&actionTabbedUnit)
	routeGenerator.routes =append(routeGenerator.Routes(), route)
	return routeGenerator
}

func (routeGenerator *RouteGenerator) Build() *core.PhpEmitterFile {
	content := routeGenerator.RouteFile.Content()


	importStmt := api.TabbedUnit(core.NewSimpleStatement(`use Illuminate\Support\Facades\Route`))
	content = append(content, &importStmt)

	fmt.Print("helloWorld")

	for _, value := range routeGenerator.RouteFile.Content() {
		fmt.Print(*value)
	}

	/*ADD ALL FUNCTION CALLS*/
	for _, functionCall := range routeGenerator.routes {
		funCall := api.TabbedUnit(functionCall)
		content = append(content, &funCall)
	}

	for _, value := range routeGenerator.RouteFile.Content() {
		fmt.Print(*value)
	}

	return routeGenerator.RouteFile
}
