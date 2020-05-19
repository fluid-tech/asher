package helper

import (
	"fmt"
	"strings"
)



func UnsignedBigInteger(colName string) string {
	return fmt.Sprintf("unsignedBigInteger('%s')", colName)
}

func BigInteger(colName string) string {
	return fmt.Sprintf("bigInteger('%s')", colName)
}

func UnsignedInteger(colName string) string {
	return fmt.Sprintf("unsignedInteger('%s')", colName)
}

func Integer(colName string) string {
	return fmt.Sprintf("integer('%s')", colName)
}

func UnsignedTinyInteger(colName string) string {
	return fmt.Sprintf("unsignedTinyInteger('%s')", colName)
}

func TinyInteger(colName string) string {
	return fmt.Sprintf("tinyInteger('%s')", colName)
}

func UnsignedMediumInteger(colName string) string {
	return fmt.Sprintf("unsignedMediumInteger('%s')", colName)
}

func MediumInteger(colName string) string {
	return fmt.Sprintf("mediumInteger('%s')", colName)
}

func String(colName string) string {
	return fmt.Sprintf("string('%s')", colName)
}

func Boolean(colName string) string {
	return fmt.Sprintf("boolean('%s')", colName)
}

func Char(colName string) string {
	return fmt.Sprintf("char('%s')", colName)
}

func Date(colName string) string {
	return fmt.Sprintf("date('%s')", colName)
}

func Double(colName string) string {
	return fmt.Sprintf("double('%s')", colName)
}

func Float(colName string) string {
	return fmt.Sprintf("float('%s')", colName)
}

func Enum(colName string, allowed []string) string {
	return fmt.Sprintf("enum('%s', %s)", colName, handleAllowedKeywordsToString(allowed))
}

func Set(colName string, allowed []string) string {
	return fmt.Sprintf("set('%s', %s)", colName, handleAllowedKeywordsToString(allowed))
}

// All Other datatype handler

func handleAllowedKeywordsToString(allowed []string) string {
	bldr := "'" + strings.Join(allowed, "', '") + "'"
	return "[" + bldr + "]"
}