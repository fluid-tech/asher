package generator

const BasicTransactor =`namespace App\Transactors;

use App\Query\CentreQuery;
use App\Transactors\Mutations\CentreMutator;

class CentreTransactor extends BaseTransactor {
    private const CLASS_NAME = 'CentreTransactor';
    public function __construct(CentreQuery $centreQuery, CentreMutator $centreMutator) {
        parent::__construct($centreQuery, $centreMutator, 'id');
        $this->className = self::CLASS_NAME;
    }


}
`