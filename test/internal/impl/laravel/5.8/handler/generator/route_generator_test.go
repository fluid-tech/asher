package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

//Route::get('/user', 'UserController@index');
func TestRouteGenerator(t *testing.T) {

	routeGenerator := generator.NewRouteGenerator()
	var table = []*api.GeneralTest{
		genRouteTest(routeGenerator, "Student", []string{}, ApiRouteFileAfterStudentWithAllRoutes),
		genRouteTest(routeGenerator, "Teacher", []string{generator.HttpGET}, ApiRouteFileAfterTeacherWithGetRoutes),
		genRouteTest(routeGenerator, "Admin",
			[]string{generator.HttpPUT, generator.HttpPOST, generator.HttpDELETE}, ApiRouteFileAfterAdminWithPATCHPOSTDELTERoutes),
	}
	api.IterateAndTest(table, t)

}

func genRouteTest(routeGen *generator.RouteGenerator, modelName string, methods []string,
	expectedOut string) *api.GeneralTest {
	return api.NewGeneralTest(routeGen.AddDefaultRestRoutes(modelName, methods).String(), expectedOut)
}
