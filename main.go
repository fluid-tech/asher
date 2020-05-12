package main

import(
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
	fmt.Println(filePath)


}
