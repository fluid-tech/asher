package generator

const EmptyMigrationWithName = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateStudentAllotmentsTable extends Migration {
    public function up() {
        Schema::create(,  function (Blueprint $table) {
}

);
    }


    public function down() {
        Schema::dropIfExists();

    }


}
`

const MigrationWithColumns = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateStudentAllotmentsTable extends Migration {
    public function up() {
        Schema::create(,  function (Blueprint $table) {
    $this->string('phone_number', 12)->unique();

    $this->string('phone_number', 12)->unique();

}

);
    }


    public function down() {
        Schema::dropIfExists();

    }


}
`
