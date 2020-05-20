package helper

import (
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
func ColTypeSwitcher(colType string, colName string, allowed []string) string {
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
		panic("not supported or wrong input in ColTypeSwitcher :- " + colType)
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
func PrimaryKeyMethodNameGenerator(colType string) string {
	switch colType {
	case "integer":
		return "increments"
	case "mediumInteger":
		return "mediumIncrements"
	case "smallInteger":
		return "smallIncrements"
	case "tinyInteger":
		return "tinyIncrements"
	case "bigInteger":
		return "bigIncrements"
	default:
		panic("Type not supported or invalid inputs")
	}
}

func UnsignedBigInteger(colName string) string {
	return normalStringDataProcessor("unsignedBigInteger", colName)
}

func BigInteger(colName string) string {
	return normalStringDataProcessor("bigInteger", colName)
}

func UnsignedInteger(colName string) string {
	return normalStringDataProcessor("unsignedInteger", colName)
}

func Integer(colName string) string {
	return normalStringDataProcessor("integer", colName)
}

func UnsignedTinyInteger(colName string) string {
	return normalStringDataProcessor("unsignedTinyInteger", colName)
}

func TinyInteger(colName string) string {
	return normalStringDataProcessor("tinyInteger", colName)
}

func UnsignedMediumInteger(colName string) string {
	return normalStringDataProcessor("unsignedMediumInteger", colName)

}

func MediumInteger(colName string) string {
	return normalStringDataProcessor("mediumInteger", colName)
}

func String(colName string, funcArgs []string) string {
	return multiParamColumnProcessor("string", colName, funcArgs)
}

func Boolean(colName string) string {
	return normalStringDataProcessor("boolean", colName)
}

func Char(colType string, dataType []string) string {
	return multiParamColumnProcessor("char", colType, dataType)
}

func Date(colName string, dataType []string) string {
	return normalStringDataProcessor("date", colName)
}

func Double(colType string, dataType []string) string {
	return multiParamColumnProcessor("double", colType, dataType)
}

func Float(colType string, dataType []string) string {
	return multiParamColumnProcessor("float", colType, dataType)
}

func Enum(colFunctionName string, allowed []string) string {
	return dataArrayProcessor(colFunctionName, allowed, "enum")
}

func Set(colFunctionName string, allowed []string) string {
	return dataArrayProcessor(colFunctionName, allowed, "set")
}

func Binary(colName string) string {
	return normalStringDataProcessor("binary", colName)
}

func DateTime(colName string, dataType []string) string {
	return multiParamColumnProcessor("dateTime", colName, nil)
}

func DateTimeTz(colName string, dataType []string) string {
	return multiParamColumnProcessor("dateTimeTz", colName, nil)
}

func Decimal(colName string, dataType []string) string {
	return multiParamColumnProcessor("dateTime", colName, nil)
}

func Geometry(colName string) string {
	return normalStringDataProcessor("geometry", colName)
}

func GeometryCollection(colName string) string {
	return normalStringDataProcessor("geometryCollection", colName)
}

func IpAddress(colName string) string {
	return normalStringDataProcessor("ipAddress", colName)
}

func Json(colName string) string {
	return normalStringDataProcessor("json", colName)
}

func Jsonb(colName string) string {
	return normalStringDataProcessor("jsonb", colName)
}

func LineString(colName string) string {
	return normalStringDataProcessor("lineString", colName)
}

func LongText(colName string) string {
	return normalStringDataProcessor("longText", colName)
}

func MacAddress(colName string) string {
	return normalStringDataProcessor("macAddress", colName)
}

func Morphs(colName string) string {
	return normalStringDataProcessor("morphs", colName)
}

func UuidMorphs(colName string) string {
	return normalStringDataProcessor("uuidMorphs", colName)
}

func MultiLineString(colName string) string {
	return normalStringDataProcessor("multiLineString", colName)
}

func MultiPoint(colName string) string {
	return normalStringDataProcessor("multiPoint", colName)
}

func MultiPolygon(colName string) string {
	return normalStringDataProcessor("multiPolygon", colName)
}

func NullableMorphs(colName string) string {
	return normalStringDataProcessor("nullableMorphs", colName)
}

func NullableUuidMorphs(colName string) string {
	return normalStringDataProcessor("nullableUuidMorphs", colName)
}

func Point(colName string) string {
	return normalStringDataProcessor("point", colName)
}

func Polygon(colName string) string {
	return normalStringDataProcessor("polygon", colName)
}

func SoftDeletes(colName string, dataType []string) string {
	return multiParamColumnProcessor("softDeletes", colName, nil)
}

func SoftDeletesTz(colName string, dataType []string) string {
	return multiParamColumnProcessor("softDeletesTz", colName, nil)
}

func Text(colName string) string {
	return normalStringDataProcessor("text", colName)
}

func Time(colName string, dataType []string) string {
	return multiParamColumnProcessor("time", colName, nil)
}

func TimeTz(colName string, dataType []string) string {
	return multiParamColumnProcessor("timeTz", colName, nil)
}

func Timestamp(colName string, dataType []string) string {
	return multiParamColumnProcessor("timestamp", colName, nil)
}

func TimestampTz(colName string, dataType []string) string {
	return multiParamColumnProcessor("timestampTz", colName, nil)
}

func Year(colName string) string {
	return normalStringDataProcessor("year", colName)
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
	return fmt.Sprintf("%s('%s', %s)", functionName, colName, handleAllowedKeywordsToString(allowed))
}

func multiParamColumnProcessor(functionName string, colName string, args []string) string {
	if len(args) > 1 {
		return fmt.Sprintf("%s('%s', %s)", functionName, colName, args[1])
	}
	return fmt.Sprintf("%s('%s')", functionName, colName)
}
