package generator

const AdminBasicQuery = `namespace App\Queries;

use App\Admin;

class AdminQuery extends BaseQuery {
    public function __construct() {
        parent::__construct("App\Admin");
    }


}
`

const TeacherBasicQuery = `namespace App\Queries;

use App\Teacher;

class TeacherQuery extends BaseQuery {
    public function __construct() {
        parent::__construct("App\Teacher");
    }


}
`


const StudentBasicQuery = `namespace App\Queries;

use App\Student;

class StudentQuery extends BaseQuery {
    public function __construct() {
        parent::__construct("App\Student");
    }


}
`
