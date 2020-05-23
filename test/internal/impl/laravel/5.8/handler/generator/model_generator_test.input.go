package generator

const ModelWithFillable = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
    protected $fillable = ["name", 
"phone_number"
];

    public static function createValidationRules() {
        return [
];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
];
    }


}
`

const EmptyModel = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
    public static function createValidationRules() {
        return [
];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
];
    }


}
`

const ModelWithHidden = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
    protected $visible = ["password", 
"gender"
];

    public static function createValidationRules() {
        return [
];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
];
    }


}
`

const ModelWithCreateValidationRules = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
    public static function createValidationRules() {
        return [
'name' => [ 'string', 'max:255', 'unique:student_allotments,name' ],
'phone_number' => [ 'string', 'max:12', 'unique:users,id' ]];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
];
    }


}
`

const ModelWithUpdateValidationRules = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
    public static function createValidationRules() {
        return [
];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
'name' => [ 'string', 'max:255', 'unique:student_allotments,name,' . $rowIds['student_allotments'] ],
'phone_number' => [ 'string', 'max:12', 'unique:users,id,' . $rowIds['users'] ]];
    }


}
`

const ModelWithUpdateValidationRulesWithoutId = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentAllotments extends Model {
    public static function createValidationRules() {
        return [
'name' => [ 'string', 'max:255', 'unique:student_allotments,name' ],
'phone_number' => [ 'string', 'max:12', 'unique:users,phone_number' ]];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
'name' => [ 'string', 'max:255', 'unique:student_allotments,name,' . $rowIds['student_allotments'] ],
'phone_number' => [ 'string', 'max:12', 'unique:users,phone_number,' . $rowIds['users'] ]];
    }


}
`
