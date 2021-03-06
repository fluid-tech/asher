package handler

import (
	"asher/internal/impl/laravel/5.8/handler"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/models"
	"asher/test/api"
	"testing"
)

func Test_Relations(t *testing.T) {
	var table = []*api.GeneralTest{
		getRelationTest("Orders", inputHasManyAllCases, inputHasOneAllCases, output1CheckForeignkeyConstraint),
		getRelationTest("Orders", inputHasManyAllCasesWithBlank1, inputHasOneAllCasesWithBlank1, output1CheckForeignkeyConstraintWithBlank1),
		getRelationTest("Orders", inputHasManyAllCasesWithBlank2, inputHasOneAllCasesWithBlank2, output1CheckForeignkeyConstraintWithBlank2),
	}
	api.IterateAndTest(table, t)
}

func getRelationTest(className string, hasMany []string, hasOne []string, expectedOut string) *api.GeneralTest {

	modelGen := generator.NewModelGenerator().SetName(className)
	context.GetFromRegistry(context.Model).AddToCtx(className, modelGen)
	handler.NewRelationshipHandler().Handle(className, models.Relation{
		HasMany: hasMany,
		HasOne:  hasOne,
	})
	return api.NewGeneralTest(modelGen.String(), expectedOut)
}
