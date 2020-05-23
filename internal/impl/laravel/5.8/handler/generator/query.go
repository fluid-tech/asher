package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
)

type QueryGenerator struct {
	api.Generator
	Class     *builder.Class
	imports   []api.TabbedUnit
	modelName string
	relation  bool
}

/**
Get the new instance of query Generator ,Query is a part of transactor pattern which handles the read operations
like getById , paginate for various frameworks like datatables.net
Parameters:
	- modelName: name of the model for which queries is to be generated
	-relation : for future use to handle nested queries
Returns:
	- instance of the generator object
Sample Usage:
	- AddDefaultRestRoutes('Order',{{Relation parameter is yet to defined }})
*/
func NewQueryGenerator(modelName string, relation bool) *QueryGenerator {
	classBuilder := builder.NewClassBuilder()
	classBuilder.SetName(modelName + "Query")
	classBuilder.SetPackage(`App\Queries`)
	return &QueryGenerator{Class: classBuilder, imports: nil, modelName: modelName, relation: relation}
}

/**
Returns the class builder object of the models query class
Returns:
	- *builder.Class
*/
func (queryGenerator *QueryGenerator) GetClass() *builder.Class {
	return queryGenerator.Class
}

/**
Builds the query class by adding imports and call to the constructor of base class (BaseQuery)
passing the fullyQualified Name as the parameter to the super call
Returns:
	- *core.Class (class object of the the query class)
*/
func (queryGenerator *QueryGenerator) Build() *core.Class {
	/*IMPORTS*/
	queryGenerator.Class.AddImport(`App\` + queryGenerator.modelName)

	/*EXTENDS*/
	queryGenerator.Class.SetExtends("BaseQuery")

	/*CONSTRUCTOR*/
	constructor := builder.NewFunctionBuilder()
	constructor.SetName("__construct").SetVisibility("public")
	fullyQualifiedModelArg := core.NewParameter(`"App\` + queryGenerator.modelName + `"`)
	callToSuperConstructor := core.NewFunctionCall("parent::__construct").AddArg(fullyQualifiedModelArg)
	constructor.AddStatement(callToSuperConstructor)
	queryGenerator.GetClass().AddFunction(constructor.GetFunction())

	return queryGenerator.Class.GetClass()
}

/**
Returns:
	- contents of Query Class in string
Sample Usage:
	-eg.input :if called for category model
	-eg.output:
	namespace App\Queries;


	use App\Category;

	class CategoryQuery extends BaseQuery
	{
		public function __construct()
		{
			parent::__construct("App\Category");
		}

	}


*/
func (queryGenerator QueryGenerator) String() string {
	queryClass := queryGenerator.Build()
	return queryClass.String()
}
