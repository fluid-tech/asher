package handler

import (
	"asher/internal/impl/laravel/5.8/handler"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/impl/laravel/5.8/handler/helper"
	"asher/test/api"
	generator2 "asher/test/internal/impl/laravel/5.8/handler/generator"
	"testing"
)

func TestAuditCol(t *testing.T) {

	var table = []*struct {
		in  []string
		out []string
	}{
		{genAuditColTest("Hello", true, true, true,
			"unsignedInteger", t), []string{generator2.ClassWithAllArgsSet, generator2.AuditColModelWithAllSet}},

		{genAuditColTest("Rnadom", false, false, true,
			"unsignedBigInteger", t), []string{generator2.ClassNoSoftDeletesAndNotTimestamp, generator2.AuditColModelWithAuditColOnly}},

		{genAuditColTest("Random", true, true, false,
			"unsignedInteger", t), []string{generator2.ClassWithSoftDeletesAndTimestamp, generator2.AuditColModelWithAuditColUnset}},

		{genAuditColTest("HelloW", false, false,
			false, "unsignedInteger", t), []string{generator2.ClassWithNoArgs, generator2.EmptyAuditCol}},
	}

	for i, element := range table {
		if element.in[0] != element.out[0] {
			t.Errorf("in test case %d expected '%s' found '%s'", i, element.out[0], element.in[0])
		}
		if element.in[1] != element.out[1] {
			t.Errorf("in test case %d expected '%s' found '%s'", i, element.out[1], element.in[1])
		}
	}

}

/**
 Returns a row indicating the following information
    - string of migration class
	- string of model class
*/
func genAuditColTest(className string, softDeletes bool, timestamp bool, auditCol bool, pkCol string, t *testing.T) []string {
	modelGen := generator.NewModelGenerator().SetName(className)
	migGen := generator.NewMigrationGenerator().SetName(className)

	context.GetFromRegistry(context.ContextMigration).AddToCtx(className, migGen)
	context.GetFromRegistry(context.ContextModel).AddToCtx(className, modelGen)

	emitterFiles, error := handler.NewAuditColHandler().Handle(className, helper.NewAuditCol(auditCol, softDeletes, timestamp, pkCol))
	if error != nil {
		t.Error(error)
	}
	if emitterFiles != nil && len(emitterFiles) > 0 {
		t.Error("audit col handler returned a file")
	}
	retrievedMigGen := api.FromContext(context.ContextMigration, className)
	if retrievedMigGen == nil {
		t.Errorf("migration for %s doesnt exist ", className)
	}
	retrievedModGen := api.FromContext(context.ContextModel, className)
	if retrievedModGen == nil {
		t.Errorf("model for %s doesnt exist ", className)
	}
	return []string{migGen.String(), modelGen.String()}
}
