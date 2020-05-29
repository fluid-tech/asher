package generator

const ApiRouteFileAfterStudentWithAllRoutes = `use Illuminate\Support\Facades\Route;
Route::post("/student", "StudentController@create");
Route::put("/student/{id}", "StudentController@update");
Route::delete("/student/{id}", "StudentController@delete");
Route::get("/student/{id}", "StudentController@findById");
Route::get("/student/all", "StudentController@getAll");
`

const ApiRouteFileAfterTeacherWithGetRoutes = `use Illuminate\Support\Facades\Route;
Route::get("/teacher/{id}", "TeacherController@findById");
Route::get("/teacher/all", "TeacherController@getAll");
`
const ApiRouteFileAfterAdminWithPATCHPOSTDELTERoutes = `use Illuminate\Support\Facades\Route;
Route::post("/admin", "AdminController@create");
Route::put("/admin/{id}", "AdminController@update");
Route::delete("/admin/{id}", "AdminController@delete");
`
