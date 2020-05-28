package builder

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
)

type CatchBlock struct {
	interfaces.CatchBlock
	catchBlock *core.CatchBlock
}

func NewCatchBlockBuilder() *CatchBlock {
	return &CatchBlock{
		catchBlock: core.NewCatchBlock(),
	}
}

func (c *CatchBlock) AddArgument(unit string) interfaces.CatchBlock {
	c.catchBlock.CatchArg = unit
	return c
}

func (c *CatchBlock) AddStatement(statement api.TabbedUnit) interfaces.CatchBlock {
	c.catchBlock.CatchStatements = append(c.catchBlock.CatchStatements, statement)
	return c
}

func (c *CatchBlock) AddStatements(statements []api.TabbedUnit) interfaces.CatchBlock {
	c.catchBlock.CatchStatements = append(c.catchBlock.CatchStatements, statements...)
	return c
}

func (c *CatchBlock) GetCatchBlock() *core.CatchBlock {
	return c.catchBlock
}
