package context

import (
	"asher/internal/impl/laravel/5.8/handler"
	"asher/internal/models"
	"testing"
)


func Test_Columns(t *testing.T) {
	//var dataTable = []struct {
	//	cols        models.Column
	//	output		string
	//}{
	//	{ models.Column{Name: "id", ColType: "bigInteger", GenerationStrategy: "auto_increment", DefaultVal: "", Table: "", Validations: "", Index: false, Allowed: nil, Invisible: false, Fillable: false, Primary: true,}, "$table->bigIncrements('id')" },
	//	{ models.Column{Name: "order_name", ColType: "string", GenerationStrategy: "", DefaultVal: "something", Table: "", Validations: "size:255|unique|nullable", Index: false, Allowed: nil, Invisible: false, Fillable: false, Primary: false,}, "$table->string('order_name)->default('something')" },
	//	{ models.Column{Name: "", ColType: "", GenerationStrategy: "", DefaultVal: "", Table: "", Validations: "", Index: false, Allowed: nil, Invisible: false, Fillable: false, Primary: false,}, "" },
	//	//{ models.Column{Name: "customer_id", ColType: "references", GenerationStrategy: "", DefaultVal: "", Table: "user:id", Validations: "exists:", Index: true, Allowed: nil, Invisible: false, Fillable: false, Primary: false,}, "" },
	//	//{  models.Column{Name: "order_colType", ColType: "enum", GenerationStrategy: "", DefaultVal: "", Table: "", Validations: "", Index: false, Allowed: ["a", "b", "c"], Invisible: true, Fillable: false, Primary: false,}},
	//}

	//{
	//Name: "", ColType: "", GenerationStrategy: "", DefaultVal: "", Table: "", Validations: "", Index: false, Allowed: nil, Hidden: false, Guarded: false, Primary: false, Unique: false, Nullable: false, OnDelete: "",
	//},
	var columnArray = []models.Column{

		{
			Name:               "id",
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
			Name:               "id_2",
			ColType:            "",
			GenerationStrategy: "uuid",
			DefaultVal:         "",
			Table:              "",
			Validations:        "unique",
			Index:              false,
			Allowed:            nil,
			Hidden:             false,
			Guarded:            true,
			Primary:            true,
			Unique:             false,
			Nullable:           false,
			OnDelete:           "",
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
			ColType:            "string",
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
			Validations:        "exists:",
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
			Allowed: 			[]string{"a", "asas", "saDASD"},
			Hidden:             false,
			Guarded:            false,
			Primary:            false,
			Unique:             true,
			Nullable:           true,
			OnDelete:           "",
		},
	}
	
	handler.NewColumnHandler().Handle("student_enrollments", columnArray)
	t.Error("Unexpected data")




	//for _, element := range classes {
	//	migration.AddToCtx(element.class.Name, element.class)
	//	if migration.GetCtx(element.expectedName).(*context.MigrationInfo).Class.Name != element.expectedName {
	//		t.Error("Unexpected data")
	//	}
	//}
}
