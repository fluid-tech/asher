package handler

const CTOrder1Controller = `namespace App\Http\Controllers;

use App\Order1;
use App\Transactors\Order1Transactor;
use App\Query\Order1Query;
use Illuminate\Http\Request;

class Order1RestController extends Controller {
    private $order1Query;
    private $order1Transactor;
    public function __construct(Order1Query $order1Query, Order1Transactor $order1Transactor) {
        $this->order1Query = $order1Query;
        $this->order1Transactor = $order1Transactor;
    }


    public function findById($id) {
        return response(['data' => $this->order1Query->findById($id)]);
    }


    public function getAll() {
        return $this->order1Query->datatables();
    }


}
`
const CTOrder1Transactor = `namespace App\Transactors;

use App\Query\Order1Query;
use App\Transactors\Mutations\Order1Mutator;

class Order1Transactor extends BaseTransactor {
    private const CLASS_NAME = 'Order1Transactor';
    public function __construct(Order1Query $order1Query, Order1Mutator $order1Mutator, $bulkDeleteColumn) {
        parent::__construct($order1Query, $order1Mutator, "id");
        $this->className = self::CLASS_NAME;
    }


}
`
const CTOrder1Mutator = `namespace App\Transactors\Mutations;

class Order1Mutator extends BaseMutator {
    public function __construct() {
        parent::__construct('App\Order1', 'id');
    }


}
`
const CTOrder1Query = `namespace App\Queries;

use App\Order1;

class Order1Query extends BaseQuery {
    public function __construct() {
        parent::__construct("App\Order1");
    }


}
`

const CTRouteFileAfterOrder1=`use Illuminate\Support\Facades\Route;
Route::get(/order1/{id}, Order1Controller@get-by-id);
Route::get(/order1/all, Order1Controller@all);
`

const CTOrder2Controller = `namespace App\Http\Controllers;

use App\Order2;
use App\Transactors\Order2Transactor;
use App\Query\Order2Query;
use Illuminate\Http\Request;

class Order2RestController extends Controller {
    private $order2Query;
    private $order2Transactor;
    public function __construct(Order2Query $order2Query, Order2Transactor $order2Transactor) {
        $this->order2Query = $order2Query;
        $this->order2Transactor = $order2Transactor;
    }


    public function findById($id) {
        return response(['data' => $this->order2Query->findById($id)]);
    }


    public function getAll() {
        return $this->order2Query->datatables();
    }


    public function create(Request $request) {
        $order2 = $this->order2Transactor->create(Auth::id(), $request->all());
        return Order2;
    }


    public function delete(Request $request, $id) {
        $order2 = $this->order2Transactor->delete($id, $request->user->id);
        return $order2;
    }


}
`
const CTOrder2Transactor = `namespace App\Transactors;

use App\Query\Order2Query;
use App\Transactors\Mutations\Order2Mutator;
use use App\Helpers\FileUploadHelper;

class Order2Transactor extends FileTransactor {
    public const IMAGE_VALIDATION_RULES = array(
        'file' => 'required|mimes:jpeg,jpg,png|max:3000'
    );
    private const CLASS_NAME = 'Order2Transactor';
    public function __construct(Order2Query $order2Query, Order2Mutator $order2Mutator, $bulkDeleteColumn) {
        parent::__construct($order2Query, $order2Mutator, "id", new FileUploadHelper(order2, self::IMAGE_VALIDATION_RULES,"png"));
        $this->className = self::CLASS_NAME;
    }


}
`
const CTOrder2Mutator = `namespace App\Transactors\Mutations;

class Order2Mutator extends BaseMutator {
    public function __construct() {
        parent::__construct('App\Order2', 'id');
    }


}
`
const CTOrder2Query = `namespace App\Queries;

use App\Order2;

class Order2Query extends BaseQuery {
    public function __construct() {
        parent::__construct("App\Order2");
    }


}
`

const CTRouteFileAfterOrder2=`use Illuminate\Support\Facades\Route;
Route::get(/order1/{id}, Order1Controller@get-by-id);
Route::get(/order1/all, Order1Controller@all);
Route::get(/order2/{id}, Order2Controller@get-by-id);
Route::get(/order2/all, Order2Controller@all);
Route::post(/order2/create, Order2Controller@create);
Route::patch(/order2/edit/{id}, Order2Controller@edit);
Route::delete(/order2/delete/{id}, Order2Controller@delete);
`

const CTOrder3Controller = `namespace App\Http\Controllers;

use App\Order3;
use App\Transactors\Order3Transactor;
use App\Query\Order3Query;
use Illuminate\Http\Request;

class Order3RestController extends Controller {
    private $order3Query;
    private $order3Transactor;
    public function __construct(Order3Query $order3Query, Order3Transactor $order3Transactor) {
        $this->order3Query = $order3Query;
        $this->order3Transactor = $order3Transactor;
    }


    public function findById($id) {
        return response(['data' => $this->order3Query->findById($id)]);
    }


    public function getAll() {
        return $this->order3Query->datatables();
    }


}
`
const CTOrder3Transactor = `namespace App\Transactors;

use App\Query\Order3Query;
use App\Transactors\Mutations\Order3Mutator;

class Order3Transactor extends ImageTransactor {
    public const IMAGE_VALIDATION_RULES = array(
        'file' => 'required|mimes:jpeg,jpg,png|max:3000'
    );
    private const CLASS_NAME = 'Order3Transactor';
    public function __construct(Order3Query $order3Query, Order3Mutator $order3Mutator, $bulkDeleteColumn, use App\Helpers\ImageUploadHelper) {
        parent::__construct($order3Query, $order3Mutator, "id", new ImageUploadHelper(order3, self::IMAGE_VALIDATION_RULES));
        $this->className = self::CLASS_NAME;
    }


}
`
const CTOrder3Mutator = `namespace App\Transactors\Mutations;

class Order3Mutator extends BaseMutator {
    public function __construct() {
        parent::__construct('App\Order3', 'id');
    }


}
`
const CTOrder3Query = `namespace App\Queries;

use App\Order3;

class Order3Query extends BaseQuery {
    public function __construct() {
        parent::__construct("App\Order3");
    }


}
`

const CTRouteFileAfterOrder3=`use Illuminate\Support\Facades\Route;
Route::get(/order1/{id}, Order1Controller@get-by-id);
Route::get(/order1/all, Order1Controller@all);
Route::get(/order2/{id}, Order2Controller@get-by-id);
Route::get(/order2/all, Order2Controller@all);
Route::post(/order2/create, Order2Controller@create);
Route::patch(/order2/edit/{id}, Order2Controller@edit);
Route::delete(/order2/delete/{id}, Order2Controller@delete);
Route::get(/order3/{id}, Order3Controller@get-by-id);
Route::get(/order3/all, Order3Controller@all);
`