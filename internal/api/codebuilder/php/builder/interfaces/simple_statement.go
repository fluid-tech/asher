package interfaces

import "asher/internal/api/codebuilder/php/core"

type SimpleStatement interface{
	SetStatement(statement string) SimpleStatement
	GetStatement() core.SimpleStatement
}
