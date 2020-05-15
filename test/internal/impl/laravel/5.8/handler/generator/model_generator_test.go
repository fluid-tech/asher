package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"fmt"
	"testing"
)

func TestFillable(t *testing.T) {
	model := generator.NewModelGenerator().SetName("student_enrollments").AddFillable("name").AddFillable("phone").Build()
	fmt.Println(model)
	t.Error("Unexpected data")
}