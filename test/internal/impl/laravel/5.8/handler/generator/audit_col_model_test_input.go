package generator

const AuditColModelWithAllSet = `namespace App;

use Illuminate\Database\Eloquent\Model;

class Hello extends Model {
    public $timestamps = true;
    use SoftDeletes;
    protected $fillable = ["created_by", 
"updated_by", 
"deleted_at"
];

    public static function createValidationRules() {
        return [
'created_by' => [ 'exists:users,id' ]];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
'deleted_at' => [ 'required', 'date_format:Y-m-d H:i:s' ],
'updated_by' => [ 'exists:users,id' ]];
    }


}
`

const AuditColModelWithAuditColOnly = `namespace App;

use Illuminate\Database\Eloquent\Model;

class Rnadom extends Model {
    protected $fillable = ["created_by", 
"updated_by"
];

    public static function createValidationRules() {
        return [
'created_by' => [ 'exists:users,id' ]];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
'updated_by' => [ 'exists:users,id' ]];
    }


}
`

const AuditColModelWithAuditColUnset = `namespace App;

use Illuminate\Database\Eloquent\Model;

class Random extends Model {
    public $timestamps = true;
    use SoftDeletes;
    protected $fillable = ["deleted_at"
];

    public static function createValidationRules() {
        return [
];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
'deleted_at' => [ 'required', 'date_format:Y-m-d H:i:s' ]];
    }


}
`

const EmptyAuditCol = `namespace App;

use Illuminate\Database\Eloquent\Model;

class HelloW extends Model {
<<<<<<< HEAD
=======
    public static function createValidationRules() {
        return [
];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
];
    }


>>>>>>> 870e47687678b640197bbdea5277941971d34423
}
`
