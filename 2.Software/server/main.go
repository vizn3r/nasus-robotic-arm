package main

import (
	"os"
	"server/app"
)

func main() {
	app.ResolveArgs(os.Args[1:])
}