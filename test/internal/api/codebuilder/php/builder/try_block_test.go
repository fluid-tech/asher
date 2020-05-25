package builder

import (
	"asher/internal/api/codebuilder/php/core"
	api2 "asher/test/api"
	"testing"
)

func TestTryBlock(t *testing.T) {

	table := []*api2.GeneralTest{
		genTryCatchBlock(true, false, BasicTryCatch),
		genTryCatchBlock(false, true, BasicTryFinally),
		genTryCatchBlock(true, true, BasicTryCatchFinally),
	}
	api2.IterateAndTest(table, t)
}

func genTryCatchBlock(catch bool, finally bool, expectedCode string) *api2.GeneralTest {
	tryBlock := core.NewTryBlock()

	tryBlock.AddStatement(core.NewSimpleStatement(`$error = 5/0`))

	if catch {
		catchBlock := core.NewCatchBlock()
		catchBlock.AddArgument("DivideByZeroException $exception")
		tryBlock.AddCatchBlock(catchBlock)
	}
	if finally {
		tryBlock.AddFinallyStatement(core.NewSimpleStatement(`echo "Hello I am in finally"`))
	}
	return api2.NewGeneralTest(tryBlock.String(), expectedCode)
}
