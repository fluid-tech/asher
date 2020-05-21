package generator

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/test/api"
	"testing"
)

func TestTransactorGeneratorTest(t *testing.T)  {
	transactorGenerator := generator.NewTransactorGenerator()
	transactorGenerator.SetIdentifier("Centre")
	//fmt.Printf(transactorGenerator.String() )
	test := api.NewGeneralTest(transactorGenerator.String(), BasicTransactor)
	api.IterateAndTest([]*api.GeneralTest{
		test,
	},t)
}
