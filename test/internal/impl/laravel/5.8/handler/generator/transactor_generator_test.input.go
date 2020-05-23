package generator

const BasicTransactor = `namespace App\Transactors;

use App\Query\OrderQuery;
use App\Transactors\Mutations\OrderMutator;

class OrderTransactor extends BaseTransactor {
    private const CLASS_NAME = 'OrderTransactor';
    public function __construct(OrderQuery $orderQuery, OrderMutator $orderMutator, $bulkDeleteColumn) {
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
    public function __construct(OrderQuery $orderQuery, OrderMutator $orderMutator, $bulkDeleteColumn) {
        parent::__construct($orderQuery, $orderMutator, "id", new FileUploadHelper(order, self::IMAGE_VALIDATION_RULES,"png"));
        $this->className = self::CLASS_NAME;
    }


}
`

const ImageTransactor = `namespace App\Transactors;

use App\Query\OrderQuery;
use App\Transactors\Mutations\OrderMutator;

class OrderTransactor extends ImageTransactor {
    public const IMAGE_VALIDATION_RULES = array(
        'file' => 'required|mimes:jpeg,jpg,png|max:3000'
    );
    private const CLASS_NAME = 'OrderTransactor';
    public function __construct(OrderQuery $orderQuery, OrderMutator $orderMutator, $bulkDeleteColumn, use App\Helpers\ImageUploadHelper) {
        parent::__construct($orderQuery, $orderMutator, "id", new ImageUploadHelper(order, self::IMAGE_VALIDATION_RULES));
        $this->className = self::CLASS_NAME;
    }


}
`