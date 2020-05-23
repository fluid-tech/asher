package generator

const BasicQuery=`namespace App\Queries;

use App\Order;

class OrderQuery extends BaseQuery {
    public function __construct() {
        parent::__construct("App\Order");
    }


}
`