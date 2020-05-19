package helper

import (
	"fmt"
	"strings"
)



func UnsignedBigInteger(colFunctionName string) string {
	return fmt.Sprintf("unsignedBigInteger('%s')", colFunctionName)
}

func BigInteger(colFunctionName string) string {
	return fmt.Sprintf("bigInteger('%s')", colFunctionName)
}

func UnsignedInteger(colFunctionName string) string {
	return fmt.Sprintf("unsignedInteger('%s')", colFunctionName)
}

func Integer(colFunctionName string) string {
	return fmt.Sprintf("integer('%s')", colFunctionName)
}

func UnsignedTinyInteger(colFunctionName string) string {
	return fmt.Sprintf("unsignedTinyInteger('%s')", colFunctionName)
}

func TinyInteger(colFunctionName string) string {
	return fmt.Sprintf("tinyInteger('%s')", colFunctionName)
}

func UnsignedMediumInteger(colFunctionName string) string {
	return fmt.Sprintf("unsignedMediumInteger('%s')", colFunctionName)
}

func MediumInteger(colFunctionName string) string {
	return fmt.Sprintf("mediumInteger('%s')", colFunctionName)
}

func String(colFunctionName string, colType string) string {
	return multiParamColumnProcessor(colType, "string")
}

func Boolean(colFunctionName string) string {
	return fmt.Sprintf("boolean('%s')", colFunctionName)
}

func Char(colFunctionName string, colType string) string {
	return multiParamColumnProcessor(colType, "char")
}

func Date(colFunctionName string) string {
	return fmt.Sprintf("date('%s')", colFunctionName)
}

func Double(colFunctionName string, colType string) string {
	return multiParamColumnProcessor(colType, "double")
}

func Float(colFunctionName string, colType string) string {
	return multiParamColumnProcessor(colType, "float")
}

func Enum(colFunctionName string, allowed []string) string {
	return dataArrayProcessor(colFunctionName, allowed, "enum")
}

func Set(colFunctionName string, allowed []string) string {
	return dataArrayProcessor(colFunctionName, allowed, "set")
}

// All Other datatype handler

func handleAllowedKeywordsToString(allowed []string) string {
	bldr := "'" + strings.Join(allowed, "', '") + "'"
	return "[" + bldr + "]"
}

func dataArrayProcessor(colFunctionName string, allowed []string, typeName string) string {
	return fmt.Sprintf("%s('%s', %s)",typeName, colFunctionName, handleAllowedKeywordsToString(allowed))
}

func multiParamColumnProcessor(colFunctionName string, typeName string) string {
	splitter := strings.Split(colFunctionName, "|")
	if len(splitter) > 1 {
		return fmt.Sprintf("%s('%s', %s)", typeName, splitter[0], splitter[1])
	}
	return fmt.Sprintf("%s('%s')", typeName, colFunctionName)
}