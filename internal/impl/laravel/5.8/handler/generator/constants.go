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
	UseSoftDeletesStr         = "use SoftDeletes"
	DefaultTimestampStr       = "public $timestamps = true"
	/*TRANSACTOR CONSTANTS*/
	ImageValidationRules = "public const IMAGE_VALIDATION_RULES =" +
		" array(\n        'file' => 'required|mimes:jpeg,jpg,png|max:3000'\n    )"
	NewImageUploadHelper     = `new ImageUploadHelper(self::BASE_PATH, self::IMAGE_VALIDATION_RULES)`
	NewBaseFileUploadHelper  = `new BaseFileUploadHelper(self::BASE_PATH, self::IMAGE_VALIDATION_RULES,"png")`
	ImageUploadHelperPath    = `App\Helpers\ImageUploadHelper`
	BaseFileUploadHelperPath = `App\Helpers\BaseFileUploadHelper`

	/*COMMON CONSTANTS*/
	VisibilityPublic    = "public"
	VisibilityPrivate   = "private"
	VisibilityProtected = "protected"
	Constructor         = "__construct"
	ParentConstructor   = "parent::__construct"

	/*Http METHODS*/
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"

	/*Http Methods Name*/
	CreateMethod   = "create"
	UpdateMethod   = "update"
	DeleteMethod   = "delete"
	GetAllMethod   = "getAll"
	FindByIdMethod = "findById"
)
