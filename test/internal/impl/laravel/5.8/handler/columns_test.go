package handler

import (
	"asher/internal/impl/laravel/5.8/handler"
	"asher/internal/models"
	"testing"
)

func TestColumns(t *testing.T) {

	var columnTestObject = []*struct {
		in  []string
		out []string
	}{
		{genColTest("StudentEnrollment", StudentEnrollmentInputArr, t), []string{ColumnTestModel, ColumnTestMigration}},
	}
	for i, element := range columnTestObject {
		if element.in[0] != element.out[0] {
			t.Errorf("in test case %d expected '%s' found '%s'", i, element.out[0], element.in[0])
		}
		if element.in[1] != element.out[1] {
			t.Errorf("in test case %d expected '%s' found '%s'", i, element.out[1], element.in[1])
		}
	}

}

func genColTest(modelName string, cols []models.Column, t *testing.T, ) []string {
	emitterFile, err := handler.NewColumnHandler().Handle(modelName, cols)

	if err != nil {
		t.Error("some errors were encountered in Handle")
	}
	if len(emitterFile) == 0 {
		t.Error("ColHandler didnt return an emitter file")
	}

	mig := fromMigReg(modelName)
	model := fromModelReg(modelName)

	if mig == nil {
		t.Errorf("migration file for %s not added to context", modelName)
	}

	if model == nil {
		t.Errorf("model file for %s not added to context", modelName)
	}

	return []string{model.String(), mig.String()}

}
