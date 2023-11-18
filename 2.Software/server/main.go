package main

import (
	"fmt"
	"os"
	"server/core"
)

func main() {

	args := core.ResolveArgs(os.Args[1:])
	fmt.Println(args)

	// t := test{Name: "test", Test: "name"}
	// var e* test = new(test)
	// fmt.Println(e) 	
	// config.ToJSON(t, "./test.json")
	// config.ParseJSON(e, "./test.json")
	// fmt.Println(e)

}