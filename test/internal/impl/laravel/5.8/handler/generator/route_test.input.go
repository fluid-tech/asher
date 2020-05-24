package generator

const ApiRouteFileAfterStudentWithAllRoutes = `use Illuminate\Support\Facades\Route;
Route::post("/student", "StudentController@create");
Route::patch("/student/{id}", "StudentController@edit");
Route::delete("/student/{id}", "StudentController@delete");
Route::get("/student/{id}", "StudentController@getById");
Route::get("/student/all", "StudentController@getAll");
`

const ApiRouteFileAfterTeacherWithGetRoutes = `use Illuminate\Support\Facades\Route;
Route::post("/student", "StudentController@create");
Route::patch("/student/{id}", "StudentController@edit");
Route::delete("/student/{id}", "StudentController@delete");
Route::get("/student/{id}", "StudentController@getById");
Route::get("/student/all", "StudentController@getAll");
Route::get("/teacher/{id}", "TeacherController@getById");
Route::get("/teacher/all", "TeacherController@getAll");
`
const ApiRouteFileAfterAdminWithPATCHPOSTDELTERoutes = `use Illuminate\Support\Facades\Route;
Route::post("/student", "StudentController@create");
Route::patch("/student/{id}", "StudentController@edit");
Route::delete("/student/{id}", "StudentController@delete");
Route::get("/student/{id}", "StudentController@getById");
Route::get("/student/all", "StudentController@getAll");
Route::get("/teacher/{id}", "TeacherController@getById");
Route::get("/teacher/all", "TeacherController@getAll");
Route::post("/admin", "AdminController@create");
Route::patch("/admin/{id}", "AdminController@edit");
Route::delete("/admin/{id}", "AdminController@delete");
`
