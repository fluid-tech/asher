package handler

import (
	"asher/internal/impl/laravel/5.8/handler"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/models"
	"testing"
)

func Test_Relations(t *testing.T) {
	modelGen := generator.NewModelGenerator().SetName("Orders")
	context.GetFromRegistry("model").AddToCtx("Orders", modelGen)
	handler.NewRelationshipHandler().Handle("Orders", models.Relation{
		HasMany: []string{"OrderProducts:order_id,pk_col", "OrderProducts", "OrderProducts:order_id"},
		HasOne:  []string{"OrderProducts:order_id,pk_col", "OrderProducts", "OrderProducts:order_id"},
	})
	stringOp := modelGen.String()
	if stringOp != output1CheckForeignkeyConstraint {
		t.Error("Test Failed Expected '" + output1CheckForeignkeyConstraint + "' But Found '" + stringOp + "'")
	}
}
