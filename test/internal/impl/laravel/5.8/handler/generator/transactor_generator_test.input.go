package generator

const StudentBasicTransactor = `namespace App\Transactors;

use App\Query\StudentQuery;
use App\Transactors\Mutations\StudentMutator;

class StudentTransactor extends BaseTransactor {
    private const CLASS_NAME = 'StudentTransactor';
    public function __construct(StudentQuery $studentQuery, StudentMutator $studentMutator) {
        parent::__construct($studentQuery, $studentMutator, "id");
        $this->className = self::CLASS_NAME;
    }


}
`

const AdminFileTransactor = `namespace App\Transactors;

use App\Query\AdminQuery;
use App\Transactors\Mutations\AdminMutator;
use use App\Helpers\BaseFileUploadHelper;

class AdminTransactor extends FileTransactor {
    public const IMAGE_VALIDATION_RULES = array(
        'file' => 'required|mimes:jpeg,jpg,png|max:3000'
    );
    private const BASE_PATH = "admin";
    private const CLASS_NAME = 'AdminTransactor';
    public function __construct(AdminQuery $adminQuery, AdminMutator $adminMutator) {
        parent::__construct($adminQuery, $adminMutator, "id", new BaseFileUploadHelper(self::BASE_PATH, self::IMAGE_VALIDATION_RULES,"png"));
        $this->className = self::CLASS_NAME;
    }


}
`

//TODO REfatcor
const TeacherImageTransactor = `namespace App\Transactors;

use App\Query\TeacherQuery;
use App\Transactors\Mutations\TeacherMutator;
use use App\Helpers\ImageUploadHelper;

class TeacherTransactor extends ImageTransactor {
    public const IMAGE_VALIDATION_RULES = array(
        'file' => 'required|mimes:jpeg,jpg,png|max:3000'
    );
    private const BASE_PATH = "teacher";
    private const CLASS_NAME = 'TeacherTransactor';
    public function __construct(TeacherQuery $teacherQuery, TeacherMutator $teacherMutator) {
        parent::__construct($teacherQuery, $teacherMutator, "id", new ImageUploadHelper(self::BASE_PATH, self::IMAGE_VALIDATION_RULES));
        $this->className = self::CLASS_NAME;
    }


}
`
