package main

import (
	"os"
	"server/app"
)

func main() {

	app.ResolveArgs(os.Args[1:])

	// t := test{Name: "test", Test: "name"}
	// var e* test = new(test)
	// fmt.Println(e) 	
	// config.ToJSON(t, "./test.json")
	// config.ParseJSON(e, "./test.json")
	// fmt.Println(e)
}