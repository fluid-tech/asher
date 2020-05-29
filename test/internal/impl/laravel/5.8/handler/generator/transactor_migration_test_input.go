package generator

const TeacherMigrationForFileURLS = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateTeacherTable extends Migration {
    public function up() {
        Schema::create('teacher',  function (Blueprint $table) {
    $table->longText('file_urls')->nullable();
}

);
    }


    public function down() {
        Schema::dropIfExists('teacher');
    }


}
`

const AdminMigrationForFileURLS = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateAdminTable extends Migration {
    public function up() {
        Schema::create('admin',  function (Blueprint $table) {
    $table->longText('file_urls')->nullable();
}

);
    }


    public function down() {
        Schema::dropIfExists('admin');
    }


}
`
