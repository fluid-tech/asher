package generator
import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)
func TestAuditColModel(t *testing.T) {
	var table = []*api.GeneralTest{
		genAuditColModelTest("Hello", true, true, true, AuditColModelWithAllSet),
		genAuditColModelTest("Hello", false, true, true, AuditColModelWithSoftDeleteUnset),
		genAuditColModelTest("Hello", true, true, false, AuditColModelWithAuditColUnset),
		genAuditColModelTest("Hello", true, false, false, AuditColModelWithAuditColAndTimestampUnset),
	}
	api.IterateAndTest(table, t)
}

func genAuditColModelTest(className string, softDeletes bool, timestamp bool, auditCol bool, expectedOut string) *api.GeneralTest {
	modelGen := generator.NewModelGenerator().SetName(className)
	generator.NewAuditColModel(modelGen).SetTimestamps(timestamp).SetAuditCol(auditCol).SetSoftDeletes(softDeletes)
	return api.NewGeneralTest(modelGen.String(), expectedOut)
}