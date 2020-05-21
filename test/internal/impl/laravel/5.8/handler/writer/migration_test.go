package writer

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/impl/laravel/5.8/handler/writer"
	"fmt"
	"testing"
)

func TestMigrationWriterHandler(t *testing.T) {
	column := core.NewSimpleStatement("$table->string('name')")
	migrationGenerator := api.Generator(generator.NewMigrationGenerator().SetName("student_allotments").
		AddColumn(*column))
	emitterFile := core.NewPhpEmitterFile("create_student_allotments",
		"/resources/projects/laravel-demo/database/migrations", migrationGenerator,
		api.Migration)
	writerHandler := writer.NewMigrationWriterHandler()
	writerHandler.BeforeHandle(emitterFile)
	bytes := writerHandler.Handle(emitterFile)
	fmt.Println(bytes)
	writerHandler.AfterHandle(emitterFile)
}
