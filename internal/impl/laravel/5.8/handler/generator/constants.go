package generator

const (
	DefaultColVal             = "unsignedBigInteger"
	CreatedByStr              = "created_by"
	UpdatedByStr              = "updated_by"
	TimestampCol              = "$table->timestamps()"
	SoftDeletesCol            = "$table->softDeletes()"
	DeletedAtStr              = "deleted_at"
	DefaultAuditColValidation = "exists:users,id"
	DeletedAtValidationRule   = "required|date_format:Y-m-d H:i:s"
	UseSoftDeletesStr		  = "use SoftDeletes"
	DefaultTimestampStr		  = "public $timestamps = true"
)
