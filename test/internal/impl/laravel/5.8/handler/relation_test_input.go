package handler

var inputHasManyAllCases []string = []string{"OrderProducts:order_id,pk_col", "OrderProducts", "OrderProducts:order_id"}
var inputHasOneAllCases []string = []string{"OrderProducts:order_id,pk_col", "OrderProducts", "OrderProducts:order_id"}

const output1CheckForeignkeyConstraint = `namespace App;

use Illuminate\Database\Eloquent\Model;

class Orders extends Model {
     function OrderProducts() {
        return $this->hasMany('App\OrderProducts','order_id','pk_col');
    }


     function OrderProducts() {
        return $this->hasMany('App\OrderProducts','orders_id','id');
    }


     function OrderProducts() {
        return $this->hasMany('App\OrderProducts','order_id','id');
    }


     function OrderProducts() {
        return $this->hasOne('App\OrderProducts','order_id','pk_col');
    }


     function OrderProducts() {
        return $this->hasOne('App\OrderProducts','orders_id','id');
    }


     function OrderProducts() {
        return $this->hasOne('App\OrderProducts','order_id','id');
    }


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

var inputHasManyAllCasesWithBlank1 []string = []string{"OrderProducts:order_id,pk_col", "OrderProducts", "OrderProducts:order_id", ""}
var inputHasOneAllCasesWithBlank1 []string = []string{"OrderProducts:order_id,pk_col", "", "OrderProducts", "OrderProducts:order_id"}

const output1CheckForeignkeyConstraintWithBlank1 = `namespace App;

use Illuminate\Database\Eloquent\Model;

class Orders extends Model {
     function OrderProducts() {
        return $this->hasMany('App\OrderProducts','order_id','pk_col');
    }


     function OrderProducts() {
        return $this->hasMany('App\OrderProducts','orders_id','id');
    }


     function OrderProducts() {
        return $this->hasMany('App\OrderProducts','order_id','id');
    }


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

var inputHasManyAllCasesWithBlank2 []string = []string{"OrderProducts:order_id,pk_col", "", "OrderProducts", "OrderProducts:order_id"}
var inputHasOneAllCasesWithBlank2 []string = []string{"OrderProducts:order_id,pk_col", "", "OrderProducts", "OrderProducts:order_id"}

const output1CheckForeignkeyConstraintWithBlank2 = `namespace App;

use Illuminate\Database\Eloquent\Model;

class Orders extends Model {
     function OrderProducts() {
        return $this->hasMany('App\OrderProducts','order_id','pk_col');
    }


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
