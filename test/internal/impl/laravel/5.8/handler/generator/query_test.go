package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"fmt"
	"testing"
)

//Route::get('/user', 'UserController@index');
func TestQueryGenerator(t *testing.T) {
	queryGenrator:=generator.NewQueryGenerator("Order", false)

	fmt.Print(queryGenrator)
}