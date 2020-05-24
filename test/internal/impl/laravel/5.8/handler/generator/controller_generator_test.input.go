package generator

const CreateRestController = `namespace App\Http\Controllers;

use App\Order;
use App\Transactors\OrderTransactor;
use App\Query\OrderQuery;
use Illuminate\Http\Request;

class OrderRestController extends Controller {
    private $orderQuery;
    private $orderTransactor;
    public function __construct(OrderQuery $orderQuery, OrderTransactor $orderTransactor) {
        $this->orderQuery = $orderQuery;
        $this->orderTransactor = $orderTransactor;
    }


    public function create(Request $request) {
        $order = $this->orderTransactor->create(Auth::id(), $request->all());
        return Order;
    }


}
`

const UpdateRestController = `namespace App\Http\Controllers;

use App\Order;
use App\Transactors\OrderTransactor;
use App\Query\OrderQuery;
use Illuminate\Http\Request;

class OrderRestController extends Controller {
    private $orderQuery;
    private $orderTransactor;
    public function __construct(OrderQuery $orderQuery, OrderTransactor $orderTransactor) {
        $this->orderQuery = $orderQuery;
        $this->orderTransactor = $orderTransactor;
    }


    public function update(Request $request) {
        $order = $this->orderTransactor->update(Auth::id(), $request->all());
        return $order;
    }


}
`

const DeleteFunctionRestController = `namespace App\Http\Controllers;

use App\Order;
use App\Transactors\OrderTransactor;
use App\Query\OrderQuery;
use Illuminate\Http\Request;

class OrderRestController extends Controller {
    private $orderQuery;
    private $orderTransactor;
    public function __construct(OrderQuery $orderQuery, OrderTransactor $orderTransactor) {
        $this->orderQuery = $orderQuery;
        $this->orderTransactor = $orderTransactor;
    }


    public function delete(Request $request, $id) {
        $order = $this->orderTransactor->delete($id, $request->user->id);
        return $order;
    }


}
`

const GetFUnctionRestController = `namespace App\Http\Controllers;

use App\Order;
use App\Transactors\OrderTransactor;
use App\Query\OrderQuery;
use Illuminate\Http\Request;

class OrderRestController extends Controller {
    private $orderQuery;
    private $orderTransactor;
    public function __construct(OrderQuery $orderQuery, OrderTransactor $orderTransactor) {
        $this->orderQuery = $orderQuery;
        $this->orderTransactor = $orderTransactor;
    }


    public function findById($id) {
        return response(['data' => $this->orderQuery->findById($id)]);
    }


    public function getAll() {
        return $this->orderQuery->datatables();
    }


}
`

const AllFunctionsRestController = `namespace App\Http\Controllers;

use App\Order;
use App\Transactors\OrderTransactor;
use App\Query\OrderQuery;
use Illuminate\Http\Request;

class OrderRestController extends Controller {
    private $orderQuery;
    private $orderTransactor;
    public function __construct(OrderQuery $orderQuery, OrderTransactor $orderTransactor) {
        $this->orderQuery = $orderQuery;
        $this->orderTransactor = $orderTransactor;
    }


    public function create(Request $request) {
        $order = $this->orderTransactor->create(Auth::id(), $request->all());
		return response()->json([
            "message"    => "Successfully Created",
            "data"       => $order
        ], 201);
    }


    public function update(Request $request) {
        $order = $this->orderTransactor->update(Auth::id(), $request->all());
        return response()->json([
            "message"    => "Successfully Updated",
            "data"       => $order
        ], 204);
    }


    public function delete(Request $request, $id) {
		$this->orderTransactor->delete($id, $request->user->id);
        return response()->json([
            "message"    => "Successfully Deleted",
        ], 204);
    }


    public function findById($id) {
		return response()->json([
			"data" => $this->orderQuery->findById($id)
        ], 200);
        return response([]);
    }


    public function getAll() {
		return response()->json([
			"data" => $this->orderQuery->paginate()
        ], 200);
    }


}
`
