package main

import (
	"bufio"
	"firmware/arm"
	"firmware/com"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	if args := os.Args; len(args) == 2 && args[1] == "t" {
		arm.ExecCode([]string{"t0"})
		return
	}

	var wg sync.WaitGroup
	fmt.Println("---------------------\n" + arm.Version + "\n---------------------\n")
	fmt.Print("See FIRMWARE.md for more info" + "\n=============================\n\n")

	arm.DocGen()

	com.CLIServer.Conf.Port = ":8080"
	// com.CLIServer.Enable()
	go com.CLIServer.StartCLI(&wg)

	s := bufio.NewScanner(os.Stdin)
	for {
		s.Scan()
		arm.ExecCode(arm.ResolveArgs(strings.Split(s.Text(), " ")))
	}
}