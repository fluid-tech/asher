package api

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/models"
	"testing"
)

type testEmitter struct {
	api.Emitter
	EmitCalled bool
	FileMap    map[string]*api.EmitterFile
}

func (t *testEmitter) Emit(value interface{}) {
	t.EmitCalled = true
}

func (t *testEmitter) GetFileMap() map[string]*api.EmitterFile {
	return t.FileMap
}

type testHandler struct{
	api.Handler
	handleCalled bool

}

func TestAsherWalker(t *testing.T) {
	var table = []struct {
		out        *testEmitter
		emitCalled bool
		className  string
		fileType   int
	}{
		{genTest("Test", 1), true, "Test", 1},
		{genTest("Hello", 2), true, "Hello", 2},
	}

	for _, element := range table {
		emittedFile := element.out.FileMap[element.className]
		if emittedFile == nil || emittedFile.FileType != element.fileType || element.out.EmitCalled != element.emitCalled || emittedFile.Klass.Name != element.className {
			t.Error("emitted file doesnt match input")
		}
	}

}

func genTest(className string, fileType int) *testEmitter {
	te := &testEmitter{
		EmitCalled: false,
		FileMap:    map[string]*api.EmitterFile{},
	}
	klass := builder.NewClassBuilder().SetName(className).GetClass()
	te.FileMap[className] = api.NewEmitterFile(className, "App/", klass, fileType)
	model := models.Model{
		Name: className,
	}
	asherObject := models.Asher{Models: []models.Model{model}}
	api.NewAsherObjectWalker(asherObject, te).Walk()
	return &testEmitter{
		EmitCalled: te.EmitCalled,
		FileMap:    te.GetFileMap(),
	}
}
