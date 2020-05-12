package main

import (
	"asher/internal"
	"fmt"
	"os"
)

func main()  {
	allArgs := os.Args
	filePath := "resources/spec-v1.json"
	if len(allArgs) > 1{
		filePath = allArgs[1]
	}else{
		fmt.Println("using default spec")
	}
	asherObject, err := internal.ToAsherObject(filePath)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("printing asher")
	fmt.Println(asherObject)

}
