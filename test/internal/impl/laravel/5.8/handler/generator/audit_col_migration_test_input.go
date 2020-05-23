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

class CreateRnadomTable extends Migration {
    public function up() {
        Schema::create('rnadom',  function (Blueprint $table) {
    $table->unsignedBigInteger('created_by');
    $table->unsignedBigInteger('updated_by')->nullable();
}

);
    }


    public function down() {
        Schema::dropIfExists('rnadom');
    }


}
`

const ClassWithSoftDeletesAndTimestamp = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateRandomTable extends Migration {
    public function up() {
        Schema::create('random',  function (Blueprint $table) {
    $table->timestamps();
    $table->softDeletes();
}

);
    }


    public function down() {
        Schema::dropIfExists('random');
    }


}
`

const ClassWithNoArgs = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateHelloWTable extends Migration {
    public function up() {
        Schema::create('hello_w',  function (Blueprint $table) {
}

);
    }


    public function down() {
        Schema::dropIfExists('hello_w');
    }


}
`
