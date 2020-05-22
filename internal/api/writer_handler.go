package api

/**
 A template for handlers that will write/generate files for the given EmitterFile.
 Every FileType must have it's own implementation for WriterHandler, so the Writer can write these type of files.
 Lifecycle:
	BeforeHandle()	->	Handle()	->	AfterHandle()
 ***********************************************************************************************************************
 Note: The methods of this handler must only be called by the Writer. To maintain its lifecycle.
*/
type WriterHandler interface {
	/**
	 A method that is called before Handler(). This can be used to perform any pre-processing operations that needs to
	 be performed on the given emitter file, if required.
	 Parameters:
		- emitterFile: instance of emitterFile that needs to be written with the required meta-data.
	*/
	BeforeHandle(emitterFile EmitterFile)

	/**
	 Handles the given emitterFile by writing the file in the given path, after performing some preprocessing if
	 required.
	 Parameters:
		- emitterFile: instance of emitterFile that needs to be written with the required meta-data.
	 Returns:
		- true if the given file was generated and written successfully.
	*/
	Handle(emitterFile EmitterFile) bool

	/**
	 A method that is called after Handler(). This can be used to perform any post-processing operations that needs to
	 be performed on the given emitter file, if required.
	 Parameters:
		- emitterFile: instance of emitterFile that needs to be written with the required meta-data.
	*/
	AfterHandle(emitterFile EmitterFile)
}
