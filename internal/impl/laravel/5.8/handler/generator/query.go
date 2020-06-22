package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
	"fmt"
	"github.com/iancoleman/strcase"
)

const QueryExtends = "BaseQuery"
const QueryNamespace = `App\Queries`

type QueryGenerator struct {
	api.Generator
	query      *core.Class
	Class      *builder.Class
	imports    []api.TabbedUnit
	identifier string
	relation   bool
}

/**
Get the new instance of query Generator ,Query is a part of transactor pattern which handles the read operations
like getById , paginate for various frameworks like datatables.net
Parameters:
	- identifier: name of the model for which queries is to be generated
	-relation : for future use to handle nested queries
Returns:
	- instance of the generator object
Sample Usage:
	- AddDefaultRestRoutes({{Relation parameter is yet to defined }})
*/
func NewQueryGenerator(relation bool) *QueryGenerator {
	classBuilder := builder.NewClassBuilder()
	return &QueryGenerator{
		query:    nil,
		Class:    classBuilder,
		imports:  nil,
		relation: relation,
	}
}

/**
Sets the identifier of the current class
Parameters:
	- identifier: string
Sample Usage:
	- SetIdentifier("ClassName")
*/
func (queryGenerator *QueryGenerator) SetIdentifier(identifier string) *QueryGenerator {
	queryGenerator.identifier = strcase.ToCamel(identifier)
	return queryGenerator
}

/**
Builds the query class by adding imports and call to the constructor of base class (BaseQuery)
passing the fullyQualified Name as the parameter to the super call
Returns:
	- *core.Class (class object of the the query class)
*/
func (queryGenerator *QueryGenerator) Build() *core.Class {
	if queryGenerator.query != nil {
		return queryGenerator.query
	}

	var className = fmt.Sprintf("%sQuery", queryGenerator.identifier)

	/*IMPORTS*/
	queryGenerator.Class.AddImport(fmt.Sprintf(ImportPathModelFmt, queryGenerator.identifier))

	/*CONSTRUCTOR*/
	constructor := builder.NewFunctionBuilder()
	constructor.SetName(FunctionNameCtor).SetVisibility(VisibilityPublic)
	fullyQualifiedModelArg := core.NewParameter(fmt.Sprintf(`"`+ImportPathModelFmt+`"`, queryGenerator.identifier))
	callToSuperConstructor := core.NewFunctionCall(FunctionNameBaseCtor).AddArg(fullyQualifiedModelArg)
	constructor.AddStatement(callToSuperConstructor)

	queryGenerator.Class.AddFunction(constructor.GetFunction()).
		SetName(className).SetPackage(QueryNamespace).SetExtends(QueryExtends)

	queryGenerator.query = queryGenerator.Class.GetClass()
	return queryGenerator.query
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
	return queryGenerator.Build().String()
}
