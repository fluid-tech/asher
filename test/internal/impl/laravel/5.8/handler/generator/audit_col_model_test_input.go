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

const AuditColModelWithAuditColOnly =`namespace App;

use Illuminate\Database\Eloquent\Model;

class Rnadom extends Model {
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

class Random extends Model {
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

const EmptyAuditCol =`namespace App;

use Illuminate\Database\Eloquent\Model;

class HelloW extends Model {
}
`


