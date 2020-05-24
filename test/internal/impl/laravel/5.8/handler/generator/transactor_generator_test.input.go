package generator

const BasicTransactor = `namespace App\Transactors;

use App\Query\OrderQuery;
use App\Transactors\Mutations\OrderMutator;

class OrderTransactor extends BaseTransactor {
    private const CLASS_NAME = 'OrderTransactor';
    public function __construct(OrderQuery $orderQuery, OrderMutator $orderMutator) {
        parent::__construct($orderQuery, $orderMutator, "id");
        $this->className = self::CLASS_NAME;
    }


}
`


const FileTransactor = `namespace App\Transactors;

use App\Query\OrderQuery;
use App\Transactors\Mutations\OrderMutator;
use use App\Helpers\FileUploadHelper;

class OrderTransactor extends FileTransactor {
    public const IMAGE_VALIDATION_RULES = array(
        'file' => 'required|mimes:jpeg,jpg,png|max:3000'
    );
    private const CLASS_NAME = 'OrderTransactor';
	private const BASE_PATH = "order"
    public function __construct(OrderQuery $orderQuery, OrderMutator $orderMutator) {
        parent::__construct($orderQuery, $orderMutator, "id", new BaseFileUploadHelper(self::BASE_PATH, self::IMAGE_VALIDATION_RULES,"png"));
        $this->className = self::CLASS_NAME;
    }


}
`

//TODO REfatcor
const ImageTransactor = `namespace App\Transactors;

use App\Query\OrderQuery;
use App\Transactors\Mutations\OrderMutator;
use App\Helpers\ImageUploadHelper

class OrderTransactor extends ImageTransactor {
    public const IMAGE_VALIDATION_RULES = array(
        'file' => 'required|mimes:jpeg,jpg,png|max:3000'
    );
    private const CLASS_NAME = 'OrderTransactor';
    public function __construct(OrderQuery $orderQuery, OrderMutator $orderMutator) {
        parent::__construct($orderQuery, $orderMutator, "id", new ImageUploadHelper(order, self::IMAGE_VALIDATION_RULES));
        $this->className = self::CLASS_NAME;
    }


}
`
