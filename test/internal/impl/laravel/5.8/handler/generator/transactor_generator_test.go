package generator

import (
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"strings"
	"testing"
)

func TestTransactorGenerator(t *testing.T) {
	var table = []*api.GeneralTest{
		genTransactorTest("Student", "base", StudentBasicTransactor),
		genTransactorTest("Admin", "file", AdminFileTransactor),
		genTransactorTest("Teacher", "image", TeacherImageTransactor),
	}
	api.IterateAndTest(table, t)
}

func genTransactorTest(modelName string, transactorType string, expectedOut string) *api.GeneralTest {
	transactorGenerator := generator.NewTransactorGenerator(transactorType).SetIdentifier(modelName)
	switch transactorType {
	case "file":
		transactorGenerator.AppendImports([]string{`App\Helpers\BaseFileUploadHelper`}).
			AddParentConstructorCallArgs(core.NewParameter(
				`new BaseFileUploadHelper(self::BASE_PATH, self::IMAGE_VALIDATION_RULES,"png")`)).
			AddTransactorMember(core.NewSimpleStatement(`private const BASE_PATH = "`+
				strings.ToLower(modelName)+`"`)).
			AddTransactorMember(core.NewSimpleStatement(
				"public const IMAGE_VALIDATION_RULES =" +
					" array(\n        'file' => 'required|mimes:jpeg,jpg,png|max:3000'\n    )"))

	case "image":
		transactorGenerator.AppendImports([]string{`App\Helpers\ImageUploadHelper`}).
			AddParentConstructorCallArgs(core.NewParameter(
				`new ImageUploadHelper(self::BASE_PATH, self::IMAGE_VALIDATION_RULES)`)).
			AddTransactorMember(core.NewSimpleStatement(`private const BASE_PATH = "`+
				strings.ToLower(modelName)+`"`)).
			AddTransactorMember(core.NewSimpleStatement(
				"public const IMAGE_VALIDATION_RULES =" +
					" array(\n        'file' => 'required|mimes:jpeg,jpg,png|max:3000'\n    )"))

	}
	//fmt.Print(transactorGenerator.String())
	return api.NewGeneralTest(transactorGenerator.String(), expectedOut)
}
