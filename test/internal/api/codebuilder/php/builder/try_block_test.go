package builder

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
	api2 "asher/test/api"
	"testing"
)

func TestTryBlock(t *testing.T) {

	table := []*api2.GeneralTest{
		genTryCatchBlockWithSingleStatement(true, false, BasicTryCatch),
		genTryCatchBlockWithSingleStatement(false, true, BasicTryFinally),
		genTryCatchBlockWithSingleStatement(true, true, BasicTryCatchFinally),
		genTryCatchBlockWithMultipleStatements(true, false, TryCatchMultipleStatement),
		genTryCatchBlockWithMultipleStatements(false, true, TryFinallyMultipleStatement),
		genTryCatchBlockWithMultipleStatements(true, true, TryCatchFinallyMultipleStatement),
	}
	api2.IterateAndTest(table, t)
}

func genTryCatchBlockWithSingleStatement(catch bool, finally bool, expectedCode string) *api2.GeneralTest {
	tryBlock := builder.NewTryBlockBuilder().AddStatement(core.NewSimpleStatement(`$error = 5/0`))

	if catch {
		tryBlock.AddCatchBlock(builder.NewCatchBlockBuilder().
			AddArgument("DivideByZeroException $exception").
			AddStatement(core.NewSimpleStatement(`echo "There was an exception"`)).GetCatchBlock())
	}
	if finally {
		tryBlock.AddFinallyStatement(core.NewSimpleStatement(`echo "Hello I am in finally"`))
	}
	return api2.NewGeneralTest(tryBlock.GetTryBlock().String(), expectedCode)
}

func genTryCatchBlockWithMultipleStatements(catch bool, finally bool, expectedCode string) *api2.GeneralTest {
	tryBlock := builder.NewTryBlockBuilder().
		AddStatements([]api.TabbedUnit{
			core.NewSimpleStatement(`$error = 5/0`),
			core.NewSimpleStatement(`echo "Hello in try block"`),
		})

	if catch {
		tryBlock.AddCatchBlock(builder.NewCatchBlockBuilder().
			AddStatements([]api.TabbedUnit{
				core.NewSimpleStatement(`echo "There was an exception"`),
				core.NewReturnStatement("$exception"),
			}).
			AddArgument("DivideByZeroException $exception").GetCatchBlock())
	}
	if finally {
		tryBlock.AddFinallyStatements([]api.TabbedUnit{
			core.NewSimpleStatement(`echo "Hello I am in finally"`),
			core.NewSimpleStatement(`echo "This block gets execute everytime"`),
		})
	}
	return api2.NewGeneralTest(tryBlock.GetTryBlock().String(), expectedCode)
}
