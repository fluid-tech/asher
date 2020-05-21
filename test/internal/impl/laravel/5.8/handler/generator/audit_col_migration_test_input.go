package generator

const ClassWithAllArgsSet = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateHelloTable extends Migration {
    public function up() {
        Schema::create('hello',  function (Blueprint $table) {
    $table->timestamps();
    $table->unsignedInteger('created_by');
    $table->unsignedInteger('updated_by')->nullable();
    $table->softDeletes();
}

);
    }


    public function down() {
        Schema::dropIfExists('hello');
    }


}
`
const ClassNoSoftDeletesAndNotTimestamp = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateHelloTable extends Migration {
    public function up() {
        Schema::create('hello',  function (Blueprint $table) {
    $table->unsignedBigInteger('created_by');
    $table->unsignedBigInteger('updated_by')->nullable();
}

);
    }


    public function down() {
        Schema::dropIfExists('hello');
    }


}
`

const ClassWithSoftDeletesAndTimestamp = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateHelloTable extends Migration {
    public function up() {
        Schema::create('hello',  function (Blueprint $table) {
    $table->timestamps();
    $table->softDeletes();
}

);
    }


    public function down() {
        Schema::dropIfExists('hello');
    }


}
`

const ClassWithNoArgs = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateHelloTable extends Migration {
    public function up() {
        Schema::create('hello',  function (Blueprint $table) {
}

);
    }


    public function down() {
        Schema::dropIfExists('hello');
    }


}
`