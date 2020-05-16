package generator

import (
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
)

func GenRouteClass(routeType int)  {
	if routeType == 0 {
		genApiRoute()
	}else if routeType == 1 {
		genWebRoute()
	}
}

func genApiRoute() {
	api
	return
}

func genWebRoute(){

}

