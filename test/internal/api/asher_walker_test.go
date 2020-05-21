package api

//
//import (
//	"asher/internal/api"
//	"asher/internal/api/codebuilder/php/builder"
//	"asher/internal/api/codebuilder/php/core"
//	"asher/internal/models"
//	"testing"
//)
//
//type testEmitter struct {
//	api.Emitter
//	EmitCalled bool
//	FileMap    map[string]*api.EmitterFile
//}
//
//func (t *testEmitter) Emit(value interface{}) {
//	t.EmitCalled = true
//}
//
//func (t *testEmitter) GetFileMap() map[string]*api.EmitterFile {
//	return t.FileMap
//}
//
//func TestAsherWalker(t *testing.T) {
//	var table = []struct {
//		out        *testEmitter
//		emitCalled bool
//		className  string
//		fileType   int
//	}{
//		{genTest("Test", 1), true, "Test", 1},
//		{genTest("Hello", 2), true, "Hello", 2},
//	}
//
//	for _, element := range table {
//		emittedFile := element.out.FileMap[element.className]
//		phpFile := (*emittedFile).(*core.PhpEmitterFile)
//		klass := (*phpFile.Content()).(*api.Generator).()
//		if emittedFile == nil || phpFile.FileType() != element.fileType || element.out.EmitCalled != element.emitCalled || klass.Name != element.className {
//			t.Error("emitted file doesnt match input")
//		}
//
//	}
//
//}
//
//func genTest(className string, fileType int) *testEmitter {
//	te := &testEmitter{
//		EmitCalled: false,
//		FileMap:    map[string]*api.EmitterFile{},
//	}
//	klass := builder.NewClassBuilder().SetName(className).GetClass()
//	klassTabbedUnit := api.TabbedUnit(klass)
//	emFile:= api.EmitterFile(core.NewPhpEmitterFile(className, "App/", []*api.TabbedUnit{&klassTabbedUnit}, fileType))
//	te.FileMap[className] = &emFile
//	model := models.Model{
//		Name: className,
//	}
//	asherObject := models.Asher{Models: []models.Model{model}}
//	api.NewAsherObjectWalker(asherObject, te).Walk()
//	return &testEmitter{
//		EmitCalled: te.EmitCalled,
//		FileMap:    te.GetFileMap(),
//	}
//}
