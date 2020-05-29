package generator

const EmptyMigrationWithName = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateStudentAllotmentsTable extends Migration {
    public function up() {
        Schema::create('student_allotments',  function (Blueprint $table) {
}

);
    }


    public function down() {
        Schema::dropIfExists('student_allotments');
    }


}
`

const StudentEmptyMigrationWithName = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateStudentTable extends Migration {
    public function up() {
        Schema::create('student',  function (Blueprint $table) {
}

);
    }


    public function down() {
        Schema::dropIfExists('student');
    }


}
`

const MigrationWithColumns = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateStudentAllotmentsTable extends Migration {
    public function up() {
        Schema::create('student_allotments',  function (Blueprint $table) {
    $table->string('name');
    $table->string('phone_number', 12)->unique();
}

);
    }


    public function down() {
        Schema::dropIfExists('student_allotments');
    }


}
`
