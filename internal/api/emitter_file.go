package api

type EmitterFile interface {
	FileName() 	string      // name of the file
	Path() 		string      // path to store it in
	Generator() 	*Generator  // The contents of the file
	FileType() 	int         // 0 - migration, 1 - model, 2- mutator, 3-transactor, 5 - controller, 6- routeFile
}

const (
	Migration  = 0
	Model      = 1
	Mutator    = 2
	Transactor = 3

	Controller = 5
	RouterFile = 6
)

const (
	MigrationPath  = `database/migrations`
	ModelPath      = `app/`
	MutatorPath    = `app/Http/Mutators`
	TransactorPath = `app/Http/Transactors`

	ControllerPath = `app/Http/Controllers`
	RouteFilePath  = `routes/`
)
