package api

type EmitterFile interface {
	FileName() string       // name of the file
	Path() string           // path to store it in
	Content() []*TabbedUnit // The contents of the file
	FileType() int          // 0 - migration, 1 - model, 2- mutator, 3-transactor, 5 - controller, 6- routeFile
}
