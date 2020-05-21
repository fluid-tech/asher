package helper

type AuditColInput struct {
	/**
	The column type specifies the action this handler performs
	For
	 	1 	- for CreatedBy,
		2 	- for UpdatedBy,
		3 	- for Both Created And UpdatedBy or formally AuditCols
		4 	- for Soft Deletes
		7 	- for CreatedBy, UpdatedBy, SoftDeletes
		8 	- for Timestamp
		15	- for CreatedBy, UpdatedBy, SoftDeletes, Timestamp
	*/
	ColumnType int
}

func NewAuditColInputFromType(auditCol bool, softDeletes bool, timestamp bool) *AuditColInput {
	return &AuditColInput{
		ColumnType: formatColumnType(auditCol, softDeletes, timestamp),
	}
}

func (input *AuditColInput) IsAuditColSet() bool {
	return input.ColumnType&3 == 3
}

func (input *AuditColInput) IsSoftDeletesSet() bool {
	return input.ColumnType&4 == 4
}

func (input *AuditColInput) IsTimestampSet() bool {
	return input.ColumnType&8 == 8
}

/**
Returns a slice containing a list of columns to be appended to the fillable array
*/
func (input *AuditColInput) GetFillableArray() []string {
	var arr []string
	if input.IsAuditColSet() {
		arr = append(arr, `"created_by"`, `"updated_by"`)
	}
	if input.IsSoftDeletesSet() {
		arr = append(arr, `"deleted_at"`)
	}
	return arr
}

func (input *AuditColInput) GetCreateValidationRules() []string {
	var arr []string
	if input.IsAuditColSet() {
		arr = append(arr, `"created_by" => "required|exists:users,id"`)
	}
	return arr
}

func (input *AuditColInput) GetUpdateValidationRules() []string {
	var arr []string
	if input.IsAuditColSet() {
		arr = append(arr, `"updated_by" => "required|exists:users,id"`)
	}
	if input.IsSoftDeletesSet() {
		// todo make this configurable
		arr = append(arr, `"deleted_at" => "required|date_format:Y-m-d H:i:s"`)
	}
	return arr
}

/**
Returns a bit mask of the booleans provided
auditCols has a value of 3
softDeletes - 4
timestamp - 8
if all are set then the integer would represent 15
*/
func formatColumnType(auditCol bool, softDeletes bool, timestamp bool) int {
	colType := 0
	if auditCol {
		colType |= 3
	}
	if softDeletes {
		colType |= 4
	}
	if timestamp {
		colType |= 8
	}
	return colType
}
