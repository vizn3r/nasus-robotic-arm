package main

import (
	"bufio"
	"cli/app"
	"fmt"
	"os"
	"strings"
)

func main() {
	if args := os.Args; len(args) != 1 {
		app.ResolveArgs(args[1:])
		os.Exit(0)
	}
	fmt.Println("Welcome to", app.Version)
	fmt.Println("\nType 'exit' to exit.\nType 'help' for help.\n\n---------------------------------------")

	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Nasus > ")
		s.Scan()
		if t := s.Text(); strings.ToLower(t) == "exit" {
			os.Exit(0)
		} else {
			app.ResolveArgs(strings.Split(t, " "))
		}
	}
}