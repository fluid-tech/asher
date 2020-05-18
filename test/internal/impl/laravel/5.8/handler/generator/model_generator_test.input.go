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
    public function createValidationRules() {
        return [
phone_number => "string|max:12|unique:users"];
    }


}
`

const ModelWithUpdateValidationRules = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
    public function updateValidationRules() {
        return [
phone_number => "string|max:12|unique:users"];
    }


}
`
