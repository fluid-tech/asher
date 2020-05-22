package writer

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/impl/laravel/5.8/handler/writer"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func BeforeTesting() {
	// Resetting the migrations
	cmd := exec.Command("php", "artisan", "migrate:reset")
	cmd.Dir = "/Users/gauravpunjabi/Coding/Projects/asher/resources/projects/laravel-demo"
	cmd.Run()

	// Deleting the migration file if it already exists..
	cmd = exec.Command("rm", "*")
	cmd.Dir = "/Users/gauravpunjabi/Coding/Projects/asher/resources/projects/laravel-demo/database/migrations/"
	cmd.Run()
}

func TestMigrationWriterHandler(t *testing.T) {
	BeforeTesting()
	const tempProjectPath = "/Users/gauravpunjabi/Coding/Projects/asher/resources/projects/laravel-demo"
	os.Chdir(tempProjectPath)

	column := core.NewSimpleStatement("$table->string('name')")
	migrationGenerator := api.Generator(generator.NewMigrationGenerator().SetName("student_allotments").
		AddColumn(*column))
	emitterFile := core.NewPhpEmitterFile("create_student_allotments",
		api.MigrationPath, migrationGenerator,
		api.Migration)
	writerHandler := writer.NewMigrationWriterHandler()
	writerHandler.BeforeHandle(emitterFile)
	bytes := writerHandler.Handle(emitterFile)
	fmt.Println(bytes)
	writerHandler.AfterHandle(emitterFile)
}
