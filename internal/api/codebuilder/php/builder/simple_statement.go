package builder

import (
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
)

type SimpleStatement struct {
	interfaces.SimpleStatement
	statement core.SimpleStatement
}

func (s *SimpleStatement) SetStatement(statement string) interfaces.SimpleStatement {
	s.statement.SimpleStatement = statement
	return s
}

func (s *SimpleStatement) GetStatement() core.SimpleStatement  {
	return s.statement
}

