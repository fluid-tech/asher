package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

//Route::get('/user', 'UserController@index');
func TestRouteGenerator(t *testing.T) {

	var table = []*api.GeneralTest{
		genRouteTest(generator.NewRouteGenerator(), "Student", []string{}, ApiRouteFileAfterStudentWithAllRoutes),
		genRouteTest(generator.NewRouteGenerator(), "Teacher", []string{generator.HttpGet}, ApiRouteFileAfterTeacherWithGetRoutes),
		genRouteTest(generator.NewRouteGenerator(), "Admin",
			[]string{generator.HttpPut, generator.HttpPost, generator.HttpDelete}, ApiRouteFileAfterAdminWithPATCHPOSTDELTERoutes),
	}
	api.IterateAndTest(table, t)

}

func genRouteTest(routeGen *generator.RouteGenerator, modelName string, methods []string,
	expectedOut string) *api.GeneralTest {
	return api.NewGeneralTest(routeGen.AddDefaultRestRoutes(modelName, methods).String(), expectedOut)
}
