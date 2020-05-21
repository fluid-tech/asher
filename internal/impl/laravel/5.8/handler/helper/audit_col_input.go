package helper

type AuditColInput struct {
	auditCol    bool
	softDeletes bool
	timestamp   bool
	PkColVal    string
}

func NewAuditCol(auditCol bool, softDeletes bool, timestamp bool, pkColVal string) *AuditColInput {
	return &AuditColInput{
		auditCol:    auditCol,
		softDeletes: softDeletes,
		timestamp:   timestamp,
		PkColVal:    pkColVal,
	}
}

/**
 Returns true if the auditCol field of this instance is set
 Returns:
	bool
 Usage:
	input.IsAuditColSet()
*/
func (input *AuditColInput) IsAuditColSet() bool {
	return input.auditCol
}

/**
 Returns true if the softDeletes field of this input is set
 Returns:
	-bool
 Usage:
	input.IsSoftDeleteSet()
*/
func (input *AuditColInput) IsSoftDeletesSet() bool {
	return input.softDeletes
}

/**
 Returns true if the timestamp field of this instance was set
 Return:
	-bool
 Usage:
	input.IsTimestampSet()
*/
func (input *AuditColInput) IsTimestampSet() bool {
	return input.timestamp
}
