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
	BasePathFmt              = `%s const BASE_PATH = "%s"`
	TransactorClassNameFmt   = `%s const CLASS_NAME = '%sTransactor'`

	VisibilityPublic     = "public"
	VisibilityPrivate    = "private"
	VisibilityProtected  = "protected"
	FunctionNameCtor     = "__construct"
	FunctionNameBaseCtor = "parent::__construct"
	QueryObjectFmt       = "%sQuery $%s"
	MutatorObjectFmt     = `%sMutator $%s`
	TransactorObjectFmt  = "%sTransactor $%s"

	/*Http METHODS*/
	HttpPost   = "POST"
	HttpGet    = "GET"
	HttpPut    = "PUT"
	HttpDelete = "DELETE"

	/*Rest Controller Methods Name*/
	MethodNameCreate   = "create"
	MethodNameUpdate   = "update"
	MethodNameDelete   = "delete"
	MethodNameGetAll   = "getAll"
	MethodNameFindById = "findById"

	ImportPathModelFmt      = `App\%s`
	ImportPathTransactorFmt = `App\Transactors\%s`
	ImportPathQueryFmt      = `App\Query\%s`
)
