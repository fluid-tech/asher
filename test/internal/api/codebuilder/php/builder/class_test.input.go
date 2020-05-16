package builder



// INITIALIZATION AREA FOR CONSTANTS
// REMEMBER TO USE SPACES INSTEAD OF TABS :(
const TestClass string = `namespace App;

use Illuminate\Database\Eloquent\Model;

class TestMutator extends BaseMutator {
    private $fullyQualifiedModel;
    public function __construct(string $fullyQualifiedModel) {
        $this->fullyQualifiedModel = $fullyQualifiedModel;

    }


}
`
const TestClass2 string = `namespace App;

use Illuminate\Database\Eloquent\Model;

class TestMutator {
    private $fullyQualifiedModel;
    public function __construct(string $fullyQualifiedModel) {
        $this->fullyQualifiedModel = $fullyQualifiedModel;

    }


}
`

const TestClass3 string = `namespace Test;

class Hello {
    private $fullyQualifiedModel;
    public function __construct(string $fullyQualifiedModel) {
        $this->fullyQualifiedModel = $fullyQualifiedModel;

    }


}
`