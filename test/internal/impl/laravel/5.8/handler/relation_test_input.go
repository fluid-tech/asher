package handler

const output_1 = `namespace App;

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
