package api

/**
 A template for handlers that will write/generate files for the given EmitterFile.
 Every FileType must have it's own implementation for WriterHandler, so the Writer can write these type of files.
 ***********************************************************************************************************************
 Note: It is recommended to create a private method named Prepare() which should be called at the start of Handle()
 method. The Prepare() will included any pre-processing operations for the given EmitterFile. If the EmitterFile doesn't
 need any pre-processing then simply no operations should be performed in it.
 */
type WriterHandler interface {
	/**
	 Handles the given emitterFile by writing the file in the given path, after performing some preprocessing if
	 required.
	 Parameters:
		- emitterFile: instance of emitterFile that needs to be written with the required meta-data.
	 Returns:
		- number of bytes that were written, if 0 bytes are returned it means the operation failed.
	 */
	Handle(emitterFile EmitterFile)		int
}