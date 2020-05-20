package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestMutatorGeneratorTest(t *testing.T)  {
	mutatorGenerator := generator.NewMutatorGenerator()
	mutatorGenerator.SetIdentifier("BatchLectureStatus")
	//fmt.Printf(mutatorGenerator.String() )
	test := api.NewGeneralTest(mutatorGenerator.String(), BasicMutator)
	api.IterateAndTest([]*api.GeneralTest{
		test,
	},t)
}
