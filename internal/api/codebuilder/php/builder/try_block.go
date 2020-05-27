package builder

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
)

type TryBlock struct {
	interfaces.TryBlock
	tryBlock *core.TryBlock
}

func NewTryBlockBuilder() *TryBlock {
	return &TryBlock{
		tryBlock: core.NewTryBlock(),
	}
}

func (t *TryBlock) AddStatement(statement api.TabbedUnit) interfaces.TryBlock {
	t.tryBlock.Statements = append(t.tryBlock.Statements, statement)
	return t
}

func (t *TryBlock) AddStatements(statements []api.TabbedUnit) interfaces.TryBlock {
	t.tryBlock.Statements = append(t.tryBlock.Statements, statements...)
	return t
}

func (t *TryBlock) AddFinallyStatement(statement api.TabbedUnit) interfaces.TryBlock {
	t.tryBlock.FinallyStatements = append(t.tryBlock.FinallyStatements, statement)
	return t
}

func (t *TryBlock) AddFinallyStatements(statements []api.TabbedUnit) interfaces.TryBlock {
	t.tryBlock.FinallyStatements = append(t.tryBlock.FinallyStatements, statements...)
	return t
}

func (t *TryBlock) AddCatchBlock(block *core.CatchBlock) interfaces.TryBlock {
	t.tryBlock.CatchBlock = append(t.tryBlock.CatchBlock, *block)
	return t
}

func (t *TryBlock) GetTryBlock() *core.TryBlock {
	return t.tryBlock
}
