package builder

import (
	"asher/internal/api/codebuilder/php/core"
	api2 "asher/test/api"
	"testing"
)

func TestForEach(t *testing.T) {
	forEach := core.NewForEach()
	forEach.AddStatement(core.NewSimpleStatement(`echo $car`))
	forEach.AddCondition("$cars as $car")
	test := api2.NewGeneralTest(forEach.String(), BasicForEach)
	api2.IterateAndTest([]*api2.GeneralTest{
		test,
	}, t)
}
