package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestAuditColModel(t *testing.T) {
	var table = []*api.GeneralTest{
		genAuditColModelTest("Hello", true, true, true, AuditColModelWithAllSet),
		genAuditColModelTest("Rnadom", false, false, true, AuditColModelWithAuditColOnly),
		genAuditColModelTest("Random", true, true, false, AuditColModelWithAuditColUnset),
		genAuditColModelTest("HelloW", false, false, false, EmptyAuditCol),
	}
	api.IterateAndTest(table, t)
}

func genAuditColModelTest(className string, softDeletes bool, timestamp bool, auditCol bool, expectedOut string) *api.GeneralTest {
	modelGen := generator.NewModelGenerator().SetName(className)
	generator.NewAuditColModel(modelGen).SetTimestamps(timestamp).SetAuditCol(auditCol).SetSoftDeletes(softDeletes)
	return api.NewGeneralTest(modelGen.String(), expectedOut)
}
