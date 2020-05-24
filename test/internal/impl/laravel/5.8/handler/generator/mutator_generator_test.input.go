package generator

const StudentBasicMutator = `namespace App\Transactors\Mutations;

class StudentMutator extends BaseMutator {
    public function __construct() {
        parent::__construct('App\Student', 'id');
    }


}
`

const AdminBasicMutator = `namespace App\Transactors\Mutations;

class AdminMutator extends BaseMutator {
    public function __construct() {
        parent::__construct('App\Admin', 'id');
    }


}
`


const TeacherBasicMutator = `namespace App\Transactors\Mutations;

class TeacherMutator extends BaseMutator {
    public function __construct() {
        parent::__construct('App\Teacher', 'id');
    }


}
`

