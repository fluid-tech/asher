package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestAuditColMigration(t *testing.T) {
	var table = []*api.GeneralTest{
		genAuditColMigrationTest("Hello", true, true, true, "unsignedInteger",
			ClassWithAllArgsSet),
		genAuditColMigrationTest("Rnadom", false, false, true, "unsignedBigInteger",
			ClassNoSoftDeletesAndNotTimestamp),
		genAuditColMigrationTest("Random", true, true, false, "unsignedInteger",
			ClassWithSoftDeletesAndTimestamp),
		genAuditColMigrationTest("HelloW", false, false, false, "unsignedInteger",
			ClassWithNoArgs),
	}

	api.IterateAndTest(table, t)

}

func genAuditColMigrationTest(className string, softDeletes bool, timestamp bool, auditCol bool, pkCol string,
	expectedOut string) *api.GeneralTest {

	migGen := generator.NewMigrationGenerator().SetName(className)
	generator.NewAuditColMigration(migGen).SetTimestamps(timestamp).SetPkCol(pkCol).
		SetAuditCols(auditCol).SetSoftDeletes(softDeletes)
	return api.NewGeneralTest(migGen.String(), expectedOut)
}
