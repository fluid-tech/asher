package core

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"testing"
)

func TestTryBlock(t *testing.T) {
	tryBlockVar := core.NewTryBlock()
	newStatement := api.TabbedUnit(core.NewSimpleStatement(
		"$this->fullyQualifiedModel = $fullyQualifiedModel"))
	tryBlockVar.Statements = append(tryBlockVar.Statements, &newStatement)
	catchBlockVar := core.NewCatchBlock()
	catchBlockVar.AddArgument("Exception $e")
	tryBlockVar.AddStatement(&newStatement)
	catchBlockVar.AddStatement(&newStatement)
	tryBlockVar.CatchBlock = []*core.CatchBlock{
		catchBlockVar.AddArgument("hello $world"),
		catchBlockVar.AddArgument("hello $world"),
		catchBlockVar.AddArgument("hello $world"),
	}
	tryBlockVar.AddFinallyStatement(&newStatement)
}
