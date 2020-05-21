package generator

const ModelWithFillable = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
	protected $fillable = ["name", 
"phone_number"
];

}
`

const EmptyModel = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
}
`

const ModelWithHidden = `namespace App;
		
use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
    protected $visible = ["password", 
"gender"
];

}
`

const ModelWithCreateValidationRules = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
    public static function createValidationRules() {
        return [
'name' => [ 'string', 'max:225', 'unique:student_allotments,name' ];
'phone_number' => [ 'string', 'max:12', 'unique:users,id' ]];
    }


}
`

const ModelWithUpdateValidationRules = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
    public static function updateValidationRules($id_rows) {
        return [
'name' => [ 'string', 'max:225', 'unique:student_allotments,name' ];
'phone_number' => [ 'string', 'max:12', 'unique:users,id' ],
    }


}
`
