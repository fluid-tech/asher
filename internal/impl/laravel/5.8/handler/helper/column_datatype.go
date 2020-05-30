package helper

import (
	"errors"
	"fmt"
	"strings"
)

/**
 This Method will return a functionName along with column name. we did this so that we can  handle multiple datatype.
 The implementation of each datatypes belongs to the helper package.
 Parameters:
	- colType: is the datatype of the column
	- colName: is the name given to the column
	- allowed: is the fixed set of values that are valid for the column
			   This is used for enum and set colTypes.
 Returns:
	- The name of laravel version of functaion name (String)
 Example:
	- primaryKeyMethodNameGenerator("integer")
*/
func ColTypeSwitcher(colType string, colName string, allowed []string) (string, error) {
	colDataType := strings.Split(colType, "|")
	switch colDataType[0] {
	// TODO : Add more column types here
	case "unsignedBigInteger":
		return UnsignedBigInteger(colName)
	case "bigInteger":
		return BigInteger(colName)
	case "unsignedInteger":
		return UnsignedInteger(colName)
	case "integer":
		return Integer(colName)
	case "unsignedTinyInteger":
		return UnsignedTinyInteger(colName)
	case "tinyInteger":
		return TinyInteger(colName)
	case "unsignedMediumInteger":
		return UnsignedMediumInteger(colName)
	case "mediumInteger":
		return MediumInteger(colName)
	case "unsignedSmallInteger":
		return UnsignedSmallInteger(colName)
	case "smallInteger":
		return smallInteger(colName)
	case "string":
		return String(colName, colDataType)
	case "boolean":
		return Boolean(colName)
	case "char":
		return Char(colName, colDataType)
	case "date":
		return Date(colName, colDataType)
	case "double":
		return Double(colName, colDataType)
	case "float":
		return Float(colName, colDataType)
	case "enum":
		return Enum(colName, allowed)
	case "set":
		return Set(colName, allowed)
	case "dateTime":
		return DateTime(colName, colDataType)
	case "dateTimeTz":
		return DateTimeTz(colName, colDataType)
	case "decimal":
		return Decimal(colName, colDataType)
	case "geometry":
		return Geometry(colName)
	case "geometryCollection":
		return GeometryCollection(colName)
	case "ipAddress":
		return IpAddress(colName)
	case "json":
		return Json(colName)
	case "jsonb":
		return Jsonb(colName)
	case "lineString":
		return LineString(colName)
	case "longText":
		return LongText(colName)
	case "macAddress":
		return MacAddress(colName)
	case "morphs":
		return Morphs(colName)
	case "uuidMorphs":
		return UuidMorphs(colName)
	case "multiLineString":
		return MultiLineString(colName)
	case "multiPoint":
		return MultiPoint(colName)
	case "multiPolygon":
		return MultiPolygon(colName)
	case "nullableMorphs":
		return NullableMorphs(colName)
	case "nullableUuidMorphs":
		return NullableUuidMorphs(colName)
	case "point":
		return Point(colName)
	case "polygon":
		return Polygon(colName)
	case "softDelete":
		return SoftDeletes(colName, colDataType)
	case "softDeleteTz":
		return SoftDeletesTz(colName, colDataType)
	case "text":
		return Text(colName)
	case "time":
		return Time(colName, colDataType)
	case "timeTz":
		return TimeTz(colName, colDataType)
	case "timestamp":
		return Timestamp(colName, colDataType)
	case "timeStampTz":
		return TimestampTz(colName, colDataType)
	case "year":
		return Year(colName)

	default:
		// TODO: Log this error and replace it with formatted error message.
		//panic("not supported or wrong input in ColTypeSwitcher :- " + colType)
		return "", errors.New("unsupported datatype")
	}
}

/**
 This Method will return a laravel version of the function name for the passed datatype Primary Key generation
 Parameters:
	- colType: Any Of the Input belonging to
			[ "integer", "mediumInteger", "smallInteger", "tinyInteger", "bigInteger" ]
 Returns:
	- The name of laravel version of functaion name (String)
 Example:
	- primaryKeyMethodNameGenerator("integer")
*/
func PrimaryKeyMethodNameGenerator(colType string) (string, error) {
	switch colType {
	case "integer":
		return "increments", nil
	case "mediumInteger":
		return "mediumIncrements", nil
	case "smallInteger":
		return "smallIncrements", nil
	case "tinyInteger":
		return "tinyIncrements", nil
	case "bigInteger":
		return "bigIncrements", nil
	default:
		return "", errors.New("type not supported or invalid inputs")
	}
}

func UnsignedBigInteger(colName string) (string, error) {
	return normalStringDataProcessor("unsignedBigInteger", colName), nil
}

func BigInteger(colName string) (string, error) {
	return normalStringDataProcessor("bigInteger", colName), nil
}

func UnsignedInteger(colName string) (string, error) {
	return normalStringDataProcessor("unsignedInteger", colName), nil
}

func Integer(colName string) (string, error) {
	return normalStringDataProcessor("integer", colName), nil
}

func UnsignedTinyInteger(colName string) (string, error) {
	return normalStringDataProcessor("unsignedTinyInteger", colName), nil
}

func TinyInteger(colName string) (string, error) {
	return normalStringDataProcessor("tinyInteger", colName), nil
}

func UnsignedMediumInteger(colName string) (string, error) {
	return normalStringDataProcessor("unsignedMediumInteger", colName), nil
}

func MediumInteger(colName string) (string, error) {
	return normalStringDataProcessor("mediumInteger", colName), nil
}

func UnsignedSmallInteger(colName string) (string, error) {
	return normalStringDataProcessor("unsignedSmallInteger", colName), nil
}

func smallInteger(colName string) (string, error) {
	return normalStringDataProcessor("smallInteger", colName), nil
}

func String(colName string, funcArgs []string) (string, error) {
	return multiParamColumnProcessor("string", colName, funcArgs), nil
}

func Boolean(colName string) (string, error) {
	return normalStringDataProcessor("boolean", colName), nil
}

func Char(colType string, dataType []string) (string, error) {
	return multiParamColumnProcessor("char", colType, dataType), nil
}

func Date(colName string, dataType []string) (string, error) {
	return normalStringDataProcessor("date", colName), nil
}

func Double(colType string, dataType []string) (string, error) {
	return multiParamColumnProcessor("double", colType, dataType), nil
}

func Float(colType string, dataType []string) (string, error) {
	return multiParamColumnProcessor("float", colType, dataType), nil
}

func Enum(colFunctionName string, allowed []string) (string, error) {
	return dataArrayProcessor(colFunctionName, allowed, "enum"), nil
}

func Set(colFunctionName string, allowed []string) (string, error) {
	return dataArrayProcessor(colFunctionName, allowed, "set"), nil
}

func Binary(colName string) string {
	return normalStringDataProcessor("binary", colName)
}

func DateTime(colName string, dataType []string) (string, error) {
	return multiParamColumnProcessor("dateTime", colName, dataType), nil
}

func DateTimeTz(colName string, dataType []string) (string, error) {
	return multiParamColumnProcessor("dateTimeTz", colName, dataType), nil
}

func Decimal(colName string, dataType []string) (string, error) {
	return multiParamColumnProcessor("decimal", colName, dataType), nil
}

func Geometry(colName string) (string, error) {
	return normalStringDataProcessor("geometry", colName), nil
}

func GeometryCollection(colName string) (string, error) {
	return normalStringDataProcessor("geometryCollection", colName), nil
}

func IpAddress(colName string) (string, error) {
	return normalStringDataProcessor("ipAddress", colName), nil
}

func Json(colName string) (string, error) {
	return normalStringDataProcessor("json", colName), nil
}

func Jsonb(colName string) (string, error) {
	return normalStringDataProcessor("jsonb", colName), nil
}

func LineString(colName string) (string, error) {
	return normalStringDataProcessor("lineString", colName), nil
}

func LongText(colName string) (string, error) {
	return normalStringDataProcessor("longText", colName), nil
}

func MacAddress(colName string) (string, error) {
	return normalStringDataProcessor("macAddress", colName), nil
}

func Morphs(colName string) (string, error) {
	return normalStringDataProcessor("morphs", colName), nil
}

func UuidMorphs(colName string) (string, error) {
	return normalStringDataProcessor("uuidMorphs", colName), nil
}

func MultiLineString(colName string) (string, error) {
	return normalStringDataProcessor("multiLineString", colName), nil
}

func MultiPoint(colName string) (string, error) {
	return normalStringDataProcessor("multiPoint", colName), nil
}

func MultiPolygon(colName string) (string, error) {
	return normalStringDataProcessor("multiPolygon", colName), nil
}

func NullableMorphs(colName string) (string, error) {
	return normalStringDataProcessor("nullableMorphs", colName), nil
}

func NullableUuidMorphs(colName string) (string, error) {
	return normalStringDataProcessor("nullableUuidMorphs", colName), nil
}

func Point(colName string) (string, error) {
	return normalStringDataProcessor("point", colName), nil
}

func Polygon(colName string) (string, error) {
	return normalStringDataProcessor("polygon", colName), nil
}

func SoftDeletes(colName string, dataType []string) (string, error) {
	return multiParamColumnProcessor("softDeletes", colName, dataType), nil
}

func SoftDeletesTz(colName string, dataType []string) (string, error) {
	return multiParamColumnProcessor("softDeletesTz", colName, dataType), nil
}

func Text(colName string) (string, error) {
	return normalStringDataProcessor("text", colName), nil
}

func Time(colName string, dataType []string) (string, error) {
	return multiParamColumnProcessor("time", colName, dataType), nil
}

func TimeTz(colName string, dataType []string) (string, error) {
	return multiParamColumnProcessor("timeTz", colName, dataType), nil
}

func Timestamp(colName string, dataType []string) (string, error) {
	return multiParamColumnProcessor("timestamp", colName, dataType), nil
}

func TimestampTz(colName string, dataType []string) (string, error) {
	return multiParamColumnProcessor("timestampTz", colName, dataType), nil
}

func Year(colName string) (string, error) {
	return normalStringDataProcessor("year", colName), nil
}

// All Other datatype processor

func handleAllowedKeywordsToString(allowed []string) string {
	bldr := "'" + strings.Join(allowed, "', '") + "'"
	return "[" + bldr + "]"
}

func normalStringDataProcessor(colFunctionName string, colName string) string {
	return fmt.Sprintf("%s('%s')", colFunctionName, colName)
}

func dataArrayProcessor(colName string, allowed []string, functionName string) string {
	if len(allowed) > 0 {
		return fmt.Sprintf("%s('%s', %s)", functionName, colName, handleAllowedKeywordsToString(allowed))
	}
	return fmt.Sprintf("%s('%s')", functionName, colName)

}

func multiParamColumnProcessor(functionName string, colName string, args []string) string {
	if len(args) > 1 {
		return fmt.Sprintf("%s('%s', %s)", functionName, colName, args[1])
	}
	return fmt.Sprintf("%s('%s')", functionName, colName)
}
