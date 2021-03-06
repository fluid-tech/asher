package main

import (
	"asher/internal"
	"asher/internal/api"
	"asher/internal/impl/laravel"
	"fmt"
	"os"
)

func main() {
	allArgs := os.Args
	filePath := "resources/spec-v1.json"
	// todo add help and other instrs
	if len(allArgs) > 1 {
		filePath = allArgs[1]
	} else {
		fmt.Println("using default spec")
	}
	asherObject, err := internal.ToAsherObject(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("printing asher")
	fmt.Println(asherObject)
	laravelEmitter := laravel.NewLaravelEmitter("5.8")
	objectWalker := api.NewAsherObjectWalker(*asherObject, *laravelEmitter)
	objectWalker.Walk()

}
