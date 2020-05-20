package laravel

import (
	"asher/internal/api"
	"fmt"
)

/**
 Laravel's implementation for writer.
 */
type Writer struct {
	api.Writer
}

/**
 Creates a new instance of this writer.
 Returns:
	- instance of a new LaravelWriter.
 */
func NewLaravelWriter() *Writer {
	return &Writer{}
}

/**
 Implementation for Walk() of base Writer. It simply walks through all the given emitterFiles and delegates them to
 their respective WriterHandlers. It also checks if each emitterFile is written successfully or not.
 Parameters:
 	- emitterFiles: a collection of emitterFiles that needs be written.
 Returns:
	- true if all the given emitterFiles are written successfully.
 */
func (writer *Writer) Walk(emitterFiles []api.EmitterFile) bool {
	status := false
	for _, emitterFile := range emitterFiles {
		writerHandler := getWriterHandler(emitterFile.FileType())
		if writerHandler != nil {
			fmt.Println("Writer found")
		} else {
			status = false
		}
	}
	return status
}

/**
 A helper method to create an instance of a WriterHandler that matches the given fileType.
 Parameters:
	- fileType: the type of WriterHandler required.
 Returns:
	- instance of WriterHandler for the given type.
 */
func getWriterHandler(fileType int) *api.WriterHandler {
	switch fileType {
	case api.Migration:
		return nil
	default:
		return nil
	}
}