package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestModelGenerator(t *testing.T) {
	var table = []*api.GeneralTest {

	}

	api.IterateAndTest(table, t)
}

func getModelWithoutFillable() *api.GeneralTest {
	generator.NewModelGenerator().SetName("S")
}


const ModelWithoutFillable string = `namespace App;

use Illuminate\Database\Eloquent\Model;

class Student extends Model {

}
`