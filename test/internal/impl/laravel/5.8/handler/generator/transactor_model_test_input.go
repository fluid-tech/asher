package generator

const TeacherModelWithFileURLS = `namespace App;

use Illuminate\Database\Eloquent\Model;

class Teacher extends Model {
    protected $fillable = ["file_urls"
];

    public static function createValidationRules() {
        return [
'file_urls' => [ 'sometimes', 'required' ],
'file_urls.urls' => [ 'array' ]];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
'file_urls' => [ 'sometimes', 'required' ],
'file_urls.urls' => [ 'array' ]];
    }


}
`

const AdminModelWithFileURLS = `namespace App;

use Illuminate\Database\Eloquent\Model;

class Admin extends Model {
    protected $fillable = ["file_urls"
];

    public static function createValidationRules() {
        return [
'file_urls' => [ 'sometimes', 'required' ],
'file_urls.urls' => [ 'array' ]];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
'file_urls' => [ 'sometimes', 'required' ],
'file_urls.urls' => [ 'array' ]];
    }


}
`
