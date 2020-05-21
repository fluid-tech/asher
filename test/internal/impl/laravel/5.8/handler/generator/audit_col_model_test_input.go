package generator

const AuditColModelWithAllSet =`namespace App;

use Illuminate\Database\Eloquent\Model;

class Hello extends Model {
    public $timestamps = true;
    use SoftDeletes;
    protected $fillable = ["created_by", 
"updated_by", 
"deleted_at"
];

    public function createValidationRules() {
        return [
"created_by" => "exists:users,id"];
    }


    public function updateValidationRules() {
        return [
"deleted_at" => "required|date_format:Y-m-d H:i:s",
"updated_by" => "exists:users,id"];
    }


}
`

const AuditColModelWithSoftDeleteUnset =`namespace App;

use Illuminate\Database\Eloquent\Model;

class Hello extends Model {
    public $timestamps = true;
    protected $fillable = ["created_by", 
"updated_by"
];

    public function createValidationRules() {
        return [
"created_by" => "exists:users,id"];
    }


    public function updateValidationRules() {
        return [
"updated_by" => "exists:users,id"];
    }


}
`

const AuditColModelWithAuditColUnset =`namespace App;

use Illuminate\Database\Eloquent\Model;

class Hello extends Model {
    public $timestamps = true;
    use SoftDeletes;
    protected $fillable = ["deleted_at"
];

    public function updateValidationRules() {
        return [
"deleted_at" => "required|date_format:Y-m-d H:i:s"];
    }


}
`

const AuditColModelWithAuditColAndTimestampUnset =`namespace App;

use Illuminate\Database\Eloquent\Model;

class Hello extends Model {
    use SoftDeletes;
    protected $fillable = ["deleted_at"
];

    public function updateValidationRules() {
        return [
"deleted_at" => "required|date_format:Y-m-d H:i:s"];
    }


}
`


