package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
)

type QueryGenerator struct {
	Class *builder.Class
	imports []*api.TabbedUnit
	modelName string
	relation bool
}

func NewQueryGenerator(modelName string, relation bool) *QueryGenerator {
	classBuilder := builder.NewClassBuilder()
	classBuilder.SetName(modelName+"Query")
	classBuilder.SetPackage(`App\Queries`)
	return &QueryGenerator{Class: classBuilder, imports: nil, modelName: modelName, relation: relation}
}

/**
Returns the array of functional calls for every model to add their routes
Returns:
	- array of *core.FunctionCall
*/
func (queryGenerator *QueryGenerator) GetClass() *builder.Class {
	return queryGenerator.Class
}

/**
Returns the array of tabbedUnits in which the imports array is followed by routes array
Returns:
	- array of tabbed units
*/
func (queryGenerator *QueryGenerator) Build() *core.Class {
	/*IMPORTS*/
	queryGenerator.Class.AddImport(`App\`+queryGenerator.modelName)

	/*EXTENDS*/
	queryGenerator.Class.SetExtends("BaseQuery")

	/*CONSTRUCTOR*/
	constructor := builder.NewFunctionBuilder()
	constructor.SetName("__construct").SetVisibility("public")
	fullyQualifiedModelArg:= api.TabbedUnit(core.NewParameter(`"App\`+queryGenerator.modelName+`"`))
	functionCall := core.NewFunctionCall("parent::__construct").AddArg(&fullyQualifiedModelArg)
	callToSuperConstructor := api.TabbedUnit(functionCall)
	constructor.AddStatement(&callToSuperConstructor)
	queryGenerator.GetClass().AddFunction(constructor.GetFunction())

	return queryGenerator.Class.GetClass()
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
func (queryGenerator *QueryGenerator) String() string {
	queryClass := queryGenerator.Build()
	return queryClass.String()
}
