package generator

const ApiRouteFileAfterOrder1 = `use Illuminate\Support\Facades\Route;
Route::get(/order1/{id}, Order1Controller@get-by-id);
Route::get(/order1/all, Order1Controller@all);
`

const CTRouteFileAfterOrder2 = `use Illuminate\Support\Facades\Route;
Route::get(/order1/{id}, Order1Controller@get-by-id);
Route::get(/order1/all, Order1Controller@all);
Route::get(/order2/{id}, Order2Controller@get-by-id);
Route::get(/order2/all, Order2Controller@all);
Route::post(/order2/create, Order2Controller@create);
Route::patch(/order2/edit/{id}, Order2Controller@edit);
Route::delete(/order2/delete/{id}, Order2Controller@delete);
`
const CTRouteFileAfterOrder3 = `use Illuminate\Support\Facades\Route;
Route::get(/order1/{id}, Order1Controller@getById);
Route::get(/order1/all, Order1Controller@getAll);
Route::get(/order2/{id}, Order2Controller@get-by-id);
Route::get(/order2/all, Order2Controller@all);
Route::post(/order2, Order2Controller@create);
Route::patch(/order2/{id}, Order2Controller@edit);
Route::delete(/order2/{id}, Order2Controller@delete);
Route::get(/order3/{id}, Order3Controller@get-by-id);
Route::get(/order3/all, Order3Controller@all);
`