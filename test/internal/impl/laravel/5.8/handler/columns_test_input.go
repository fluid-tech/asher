package handler

import (
	"asher/internal/models"
)

const test_1_tableName = "student_enrollments"
var test_1_columnInputArray = []models.Column{

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
		Primary:            true,
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
		Primary:            true,
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
		Primary:            true,
		Unique:             false,
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
		Primary:            true,
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

var test_1_fillableExpectedOutput = []string{`"id_int"`, `"id_medium"`, `"id_small"`, `"id_tiny"`, `"id_big"`, `"id_uuid"`, `"order_id"`, `"order_id"`, `"order_id"`}
var test_1_hiddenExpectedOutput = []string{`"id_int"`, `"id_medium"`, `"id_small"`, `"id_tiny"`, `"id_big"`, `"id_uuid"`, `"order_id"`, `"order_id"`, `"order_id"`, `"name"`, `"dummy_column"`}


const migration_output_up = `public function up() {
    Schema::create('student_enrollments',  function (Blueprint $table) {
    $table->increments('id_int');
    $table->mediumIncrements('id_medium');
    $table->smallIncrements('id_small');
    $table->tinyIncrements('id_tiny');
    $table->bigIncrements('id_big');
    $table->uuid('id_uuid')->primary();
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

`

const migration_output_down = `public function down() {
    Schema::dropIfExists('student_enrollments');
}

`