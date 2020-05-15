package generator

import "asher/internal/api/codebuilder/php/builder"

type MigrationGenerator struct {
	builder.Class
}

func NewMigrationGenerator() *MigrationGenerator {
	return &MigrationGenerator{}
}

func (migrationGenerator *MigrationGenerator) GenerateMigrationTemplate(migrationClassName string) {
	migrationGenerator.Class.AddImports([]string{
		"use Illuminate\\Database\\Migrations\\Migration",
		"use Illuminate\\Database\\Schema\\Blueprint",
		"use Illuminate\\Support\\Facades\\Schema",
	})

	migrationGenerator.Class.SetExtends("Migration")

	upFunction := builder.Function.SetName("up")
	upFunction.SetVisibility("public")
	upFunction.AddStatement("Schema::create('users', function (Blueprint $table)")



	migrationGenerator.Class.A
}