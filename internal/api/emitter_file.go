package api

type EmitterFile interface {
	FileName() string      // name of the file
	Path() string          // path to store it in
	Generator() *Generator // The contents of the file
	FileType() int         // 0 - migration, 1 - model, 2- mutator, 3-transactor, 5 - controller, 6- routeFile
}

const (
	MIGRATION  = 0
	MODEL      = 1
	MUTATOR    = 2
	TRANSACTOR = 3

	CONTROLLER = 5
	ROUTEFILE  = 6
)

const (
	MIGRATION_PATH  = ""
	MODEL_PATH      = ""
	MUTATOR_PATH    = ""
	TRANSACTOR_PATH = ""

	CONTROLLER_PATH = ""
	ROUTEFILE_PATH  = ""
)