package generator

const CreateRestController = `namespace App\Http\Controllers;

use App\Order;
use App\Transactors\OrderTransactor;
use App\Query\OrderQuery;
use Illuminate\Http\Request;
use App\Helpers\ResponseHelper;

class OrderRestController extends Controller {
    private $orderQuery;
    private $orderTransactor;
    public function __construct(OrderQuery $orderQuery, OrderTransactor $orderTransactor) {
        $this->orderQuery = $orderQuery;
        $this->orderTransactor = $orderTransactor;
    }


    public function create(Request $request) {
        $order = $this->orderTransactor->create(Auth::id(), $request->all());
        return ResponseHelper::create($order);
    }


}
`

const UpdateRestController = `namespace App\Http\Controllers;

use App\Order;
use App\Transactors\OrderTransactor;
use App\Query\OrderQuery;
use Illuminate\Http\Request;
use App\Helpers\ResponseHelper;

class OrderRestController extends Controller {
    private $orderQuery;
    private $orderTransactor;
    public function __construct(OrderQuery $orderQuery, OrderTransactor $orderTransactor) {
        $this->orderQuery = $orderQuery;
        $this->orderTransactor = $orderTransactor;
    }


    public function update(Request $request) {
        $order = $this->orderTransactor->update(Auth::id(), $request->all());
        return ResponseHelper::update($order);
    }


}
`

const DeleteFunctionRestController = `namespace App\Http\Controllers;

use App\Order;
use App\Transactors\OrderTransactor;
use App\Query\OrderQuery;
use Illuminate\Http\Request;
use App\Helpers\ResponseHelper;

class OrderRestController extends Controller {
    private $orderQuery;
    private $orderTransactor;
    public function __construct(OrderQuery $orderQuery, OrderTransactor $orderTransactor) {
        $this->orderQuery = $orderQuery;
        $this->orderTransactor = $orderTransactor;
    }


    public function delete(Request $request, $id) {
        $order = $this->orderTransactor->delete($id, $request->user->id);
        return ResponseHelper::delete($order);
    }


}
`

const GetFUnctionRestController = `namespace App\Http\Controllers;

use App\Order;
use App\Transactors\OrderTransactor;
use App\Query\OrderQuery;
use Illuminate\Http\Request;
use App\Helpers\ResponseHelper;

class OrderRestController extends Controller {
    private $orderQuery;
    private $orderTransactor;
    public function __construct(OrderQuery $orderQuery, OrderTransactor $orderTransactor) {
        $this->orderQuery = $orderQuery;
        $this->orderTransactor = $orderTransactor;
    }


    public function findById($id) {
        return response()->json(['data' => $this->orderQuery->findById($id)], 200);
    }


    public function getAll() {
        return response()->json(['data' => $this->orderQuery->paginate()], 200);
    }


}
`

const AllFunctionsRestController = `namespace App\Http\Controllers;

use App\Order;
use App\Transactors\OrderTransactor;
use App\Query\OrderQuery;
use Illuminate\Http\Request;
use App\Helpers\ResponseHelper;

class OrderRestController extends Controller {
    private $orderQuery;
    private $orderTransactor;
    public function __construct(OrderQuery $orderQuery, OrderTransactor $orderTransactor) {
        $this->orderQuery = $orderQuery;
        $this->orderTransactor = $orderTransactor;
    }


    public function create(Request $request) {
        $order = $this->orderTransactor->create(Auth::id(), $request->all());
        return ResponseHelper::create($order);
    }


    public function update(Request $request) {
        $order = $this->orderTransactor->update(Auth::id(), $request->all());
        return ResponseHelper::update($order);
    }


    public function delete(Request $request, $id) {
        $order = $this->orderTransactor->delete($id, $request->user->id);
        return ResponseHelper::delete($order);
    }


    public function findById($id) {
        return response()->json(['data' => $this->orderQuery->findById($id)], 200);
    }


    public function getAll() {
        return response()->json(['data' => $this->orderQuery->paginate()], 200);
    }


}
`

const StudentController = `namespace App\Http\Controllers;

use App\Student;
use App\Transactors\StudentTransactor;
use App\Query\StudentQuery;
use Illuminate\Http\Request;
use App\Helpers\ResponseHelper;

class StudentRestController extends Controller {
    private $studentQuery;
    private $studentTransactor;
    public function __construct(StudentQuery $studentQuery, StudentTransactor $studentTransactor) {
        $this->studentQuery = $studentQuery;
        $this->studentTransactor = $studentTransactor;
    }


    public function create(Request $request) {
        $student = $this->studentTransactor->create(Auth::id(), $request->all());
        return ResponseHelper::create($student);
    }


    public function update(Request $request) {
        $student = $this->studentTransactor->update(Auth::id(), $request->all());
        return ResponseHelper::update($student);
    }


    public function delete(Request $request, $id) {
        $student = $this->studentTransactor->delete($id, $request->user->id);
        return ResponseHelper::delete($student);
    }


    public function findById($id) {
        return response()->json(['data' => $this->studentQuery->findById($id)], 200);
    }


    public function getAll() {
        return response()->json(['data' => $this->studentQuery->paginate()], 200);
    }


}
`

const TeacherController = `namespace App\Http\Controllers;

use App\Teacher;
use App\Transactors\TeacherTransactor;
use App\Query\TeacherQuery;
use Illuminate\Http\Request;
use App\Helpers\ResponseHelper;

class TeacherRestController extends Controller {
    private $teacherQuery;
    private $teacherTransactor;
    public function __construct(TeacherQuery $teacherQuery, TeacherTransactor $teacherTransactor) {
        $this->teacherQuery = $teacherQuery;
        $this->teacherTransactor = $teacherTransactor;
    }


    public function findById($id) {
        return response()->json(['data' => $this->teacherQuery->findById($id)], 200);
    }


    public function getAll() {
        return response()->json(['data' => $this->teacherQuery->paginate()], 200);
    }


}
`

const AdminController = `namespace App\Http\Controllers;

use App\Admin;
use App\Transactors\AdminTransactor;
use App\Query\AdminQuery;
use Illuminate\Http\Request;
use App\Helpers\ResponseHelper;

class AdminRestController extends Controller {
    private $adminQuery;
    private $adminTransactor;
    public function __construct(AdminQuery $adminQuery, AdminTransactor $adminTransactor) {
        $this->adminQuery = $adminQuery;
        $this->adminTransactor = $adminTransactor;
    }


    public function delete(Request $request, $id) {
        $admin = $this->adminTransactor->delete($id, $request->user->id);
        return ResponseHelper::delete($admin);
    }


}
`
