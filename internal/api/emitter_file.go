package api

type EmitterFile struct {
	FileName string // name of the file
	Path     string // path to store it in
	Content  string // content of the file
	FileType int	// 0 - migration, 1 - model, 2- mutator, 3-transactor
}

func NewEmitterFile(fileName string, path string, content string, fileType int) *EmitterFile {
	return &EmitterFile{
		FileName:fileName,
		Path:     path,
		Content:  content,
		FileType: fileType,
	}
}
