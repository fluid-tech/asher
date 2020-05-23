package generator

const BasicMutator = `namespace App\Transactors\Mutations;

class BatchLectureStatusMutator extends BaseMutator {
    public function __construct() {
        parent::__construct('App\BatchLectureStatus', 'id');
    }


}
`
