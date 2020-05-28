package handler

import (
	"asher/internal/models"
)

const TableStudentEnrollments = "student_enrollments"

var StudentEnrollmentInputArr = []models.Column{
	{
		Name:               "id_int",
		ColType:            "integer",
		GenerationStrategy: "auto_increment",
		DefaultVal:         "",
		Table:              "",
		Validations:        "min:100|max:334|number|unique:order,order_id",
		Index:              false,
		Allowed:            nil,
		Hidden:             true,
		Guarded:            true,
		Primary:            true,
		Unique:             false,
		Nullable:           false,
		OnDelete:           "",
	},
	{
		Name:               "id_medium",
		ColType:            "mediumInteger",
		GenerationStrategy: "auto_increment",
		DefaultVal:         "",
		Table:              "",
		Validations:        "max:334|unique|min:100|min:100|number",
		Index:              false,
		Allowed:            nil,
		Hidden:             true,
		Guarded:            true,
		Primary:            false,
		Unique:             false,
		Nullable:           false,
		OnDelete:           "",
	},
	{
		Name:               "id_small",
		ColType:            "smallInteger",
		GenerationStrategy: "auto_increment",
		DefaultVal:         "",
		Table:              "",
		Validations:        "unique:orders",
		Index:              false,
		Allowed:            nil,
		Hidden:             true,
		Guarded:            true,
		Primary:            false,
		Unique:             false,
		Nullable:           false,
		OnDelete:           "",
	},
	{
		Name:               "id_tiny",
		ColType:            "tinyInteger",
		GenerationStrategy: "auto_increment",
		DefaultVal:         "",
		Table:              "",
		Validations:        "unique",
		Index:              false,
		Allowed:            nil,
		Hidden:             true,
		Guarded:            true,
		Primary:            false,
		Unique:             true,
		Nullable:           false,
		OnDelete:           "",
	},
	{
		Name:               "id_big",
		ColType:            "bigInteger",
		GenerationStrategy: "auto_increment",
		DefaultVal:         "",
		Table:              "",
		Validations:        "unique",
		Index:              false,
		Allowed:            nil,
		Hidden:             true,
		Guarded:            true,
		Primary:            true,
		Unique:             false,
		Nullable:           false,
		OnDelete:           "",
	},
	{
		Name:               "id_uuid",
		ColType:            "",
		GenerationStrategy: "uuid",
		DefaultVal:         "",
		Table:              "",
		Validations:        "",
		Index:              false,
		Allowed:            nil,
		Hidden:             true,
		Guarded:            true,
		Primary:            false,
		Unique:             false,
		Nullable:           false,
		OnDelete:           "",
	},
	{
		Name:               "order_id",
		ColType:            "reference",
		GenerationStrategy: "",
		DefaultVal:         "",
		Table:              "Orders:id",
		Validations:        "",
		Index:              false,
		Allowed:            nil,
		Hidden:             true,
		Guarded:            true,
		Primary:            false,
		Unique:             false,
		Nullable:           false,
		OnDelete:           "cascade",
	},
	{
		Name:               "order_id",
		ColType:            "reference",
		GenerationStrategy: "",
		DefaultVal:         "",
		Table:              "Orders:id",
		Validations:        "",
		Index:              false,
		Allowed:            nil,
		Hidden:             true,
		Guarded:            true,
		Primary:            false,
		Unique:             false,
		Nullable:           false,
		OnDelete:           "set null",
	},
	{
		Name:               "order_id",
		ColType:            "reference",
		GenerationStrategy: "",
		DefaultVal:         "",
		Table:              "Orders:id",
		Validations:        "",
		Index:              false,
		Allowed:            nil,
		Hidden:             true,
		Guarded:            true,
		Primary:            false,
		Unique:             false,
		Nullable:           true,
		OnDelete:           "set null",
	},
	{
		Name:               "name",
		ColType:            "string",
		GenerationStrategy: "",
		DefaultVal:         "",
		Table:              "",
		Validations:        "",
		Index:              false,
		Allowed:            nil,
		Hidden:             true,
		Guarded:            false,
		Primary:            false,
		Unique:             false,
		Nullable:           false,
		OnDelete:           "",
	},
	{
		Name:               "description",
		ColType:            "string|255",
		GenerationStrategy: "",
		DefaultVal:         "default description",
		Table:              "",
		Validations:        "",
		Index:              false,
		Allowed:            nil,
		Hidden:             false,
		Guarded:            false,
		Primary:            false,
		Unique:             true,
		Nullable:           true,
		OnDelete:           "",
	},
	{
		Name:               "order_id",
		ColType:            "bigInteger",
		GenerationStrategy: "",
		DefaultVal:         "0",
		Table:              "orders:id",
		Validations:        "",
		Index:              true,
		Allowed:            nil,
		Hidden:             false,
		Guarded:            false,
		Primary:            false,
		Unique:             true,
		Nullable:           true,
		OnDelete:           "cascade",
	},
	{
		Name:               "dummy_column",
		ColType:            "enum",
		GenerationStrategy: "",
		DefaultVal:         "",
		Table:              "",
		Validations:        "",
		Index:              false,
		Allowed:            []string{"a", "asas", "saDASD"},
		Hidden:             true,
		Guarded:            false,
		Primary:            false,
		Unique:             true,
		Nullable:           true,
		OnDelete:           "",
	},
}

const ColumnTestMigration = `use Illuminate\Database\Migrations\Migration;
use Illuminate\DatabaseSchema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateStudentEnrollmentTable extends Migration {
    public function up() {
        Schema::create('student_enrollment',  function (Blueprint $table) {
    $table->increments('id_int');
    $table->mediumInteger('id_medium');
    $table->unsupported datatype;
    $table->tinyInteger('id_tiny')->unique();
    $table->bigIncrements('id_big');
    $table->unsupported datatype;
    $table->foreign('order_id')->references('id')->on('Orders')->onDelete('cascade');
    $table->foreign('order_id')->references('id')->on('Orders')->onDelete('set null');
    $table->foreign('order_id')->references('id')->on('Orders')->onDelete('set null')->nullable();
    $table->string('name');
    $table->string('description', 255)->default('default description')->nullable()->unique();
    $table->bigInteger('order_id')->default('0')->nullable()->unique()->index();
    $table->enum('dummy_column', ['a', 'asas', 'saDASD'])->nullable()->unique();
}

);
    }


    public function down() {
        Schema::dropIfExists('student_enrollment');
    }


}
`
const ColumnTestModel = `namespace App;

use Illuminate\Database\Eloquent\Model;

class StudentEnrollment extends Model {
    protected $fillable = ["id_int", 
"id_medium", 
"id_small", 
"id_tiny", 
"id_big", 
"id_uuid", 
"order_id", 
"order_id", 
"order_id"
];

    protected $visible = ["id_int", 
"id_medium", 
"id_small", 
"id_tiny", 
"id_big", 
"id_uuid", 
"order_id", 
"order_id", 
"order_id", 
"name", 
"dummy_column"
];

    public static function createValidationRules() {
        return [
'id_big' => [ 'unique:StudentEnrollment,id_big' ],
'id_int' => [ 'min:100', 'max:334', 'number', 'unique:order,order_id' ],
'id_medium' => [ 'max:334', 'unique:StudentEnrollment,id_medium', 'min:100', 'min:100', 'number' ],
'id_small' => [ 'unique:orders,id_small' ],
'id_tiny' => [ 'unique:StudentEnrollment,id_tiny' ]];
    }


    public static function updateValidationRules(array $rowIds) {
        return [
'id_big' => [ 'unique:StudentEnrollment,id_big,' . $rowIds['StudentEnrollment'] ],
'id_int' => [ 'min:100', 'max:334', 'number', 'unique:order,order_id,' . $rowIds['order'] ],
'id_medium' => [ 'max:334', 'unique:StudentEnrollment,id_medium,' . $rowIds['StudentEnrollment'], 'min:100', 'min:100', 'number' ],
'id_small' => [ 'unique:orders,id_small,' . $rowIds['orders'] ],
'id_tiny' => [ 'unique:StudentEnrollment,id_tiny,' . $rowIds['StudentEnrollment'] ]];
    }


}
`
