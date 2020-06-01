package builder

const Ctor = `public static function __construct(BaseMutator $mutator, BaseQuery $query) {
    $this->query = $query;
    $this->mutator = $mutator;
}

`
const Ctor2 = `public static function __construct(BaseMutator $mutator, BaseQuery $query, ImageHandler $imageHandler) {
    $this->query = $query;
    $this->mutator = $mutator;
}

`

const TestFunction = `protected function up($hello, $world) {
    return $world+$hello;
}

`
