package app

import (
	"fmt"
	"server/arm"
	"server/com"
	"server/util"
	"strconv"
	"strings"
	"sync"
)

var Version = "Nasus Firmware v0.0.1" // VERSION OF SOFTWARE

type Command struct {
	name   string
	desc   string
	call   []string
	run    func([]string)
	params int
}

// Make sure that args will not be empty so the program won't throw ugly errors :)
func resolveArgs(args []string, params int) []string {
	if len(args) >= params {
		return args
	}
	return []string{"", "", "", "", "", "", "", ""}
}

// Find and return argument based on input if it exixst, otherwise return empty argument with no name
func findCommand(cmd string, cmds []Command) Command {
	for _, c := range cmds {
		if strings.Compare(c.name, cmd) == 0 || util.StringArrayHas(c.call, cmd) {
			return c
		} 
	}
	fmt.Println("Argument not found.")
	helpFunc(cmds)
	return Command{name: ""}
}

// Help menu
func helpFunc(args []Command) {
	out := Version + "\nby vizn3r\n\n\nArguments:\n"
	for _, a := range args {
		out += "	" + strings.Join(append([]string{a.name}, a.call...), " / ") + " " + a.desc + "\n"
	}
	fmt.Println(out)
}

// Resolves user arguments
func ResolveArgs(args []string) {

	// Application arguments
	var cmds = make([]Command, 0)
	cmds = []Command{
		{name: "\nTEST: "},
		{
			name: "test",
			desc: "[...args] - test command",
			call: []string{"t", "tst"},
			run: func(s []string) {
				fmt.Println(s)
			},
		},

		{name: "\nMATH: "},
		{
			name: "degtostep",
			desc: "[angles, dest, step] - Convert degrees to steps",
			call: []string{"dts"},
			params: 3,
			run: func(s []string) {
				s0, _ := strconv.ParseFloat(s[0], 64)
				s1, _ := strconv.ParseFloat(s[1], 64)
				s2, _ := strconv.ParseFloat(s[2], 64)
				fmt.Println(arm.AngleToSteps(s0, s1, s2))
			},
		},
		{
			name: "steptodeg",
			desc: "[angle, steps, step] - Convert steps to degrees",
			call: []string{"std"},
			params: 3,
			run: func(s []string) {
				s0, _ := strconv.ParseFloat(s[0], 64)
				s1, _ := strconv.ParseFloat(s[1], 64)
				s2, _ := strconv.ParseFloat(s[2], 64)
				fmt.Println(arm.StepsToAngle(s0, s1, s2))
			},
		},

		{name: "\nARM CONTROL: "},
		{
			name: "code",
			desc: "[code, ...args] | [filePath] - Execute arm code",
			call: []string{"c"},
			params: 1,
			run: func(s []string) {
				arm.ResolveCode(s)
			},
		},

		{name: "\nAPPLICATION"},
		{
			name: "http",
			desc: "- Start HTTP server",
			run: func(s []string) {
				c := ConfigFromFile()
				var wg sync.WaitGroup
				wg.Add(1)	
				go func () {
					com.StartHTTP(c.HTTP)
					defer wg.Done()
				}()
				wg.Wait()
			},
		},
		{
			name: "btconnect",
			call: []string{"btc"},
			desc: "- Connect to bluetooth device",
			run: func(s []string) {
				com.ConnectBT()
			},
		},
		{
			name: "listen",
			call: []string{"l"},
			desc: "[baud] - Print serial data from port",
			run: func(s []string) {
				for {com.ReadSerial()}
			},
		},
		{
			name: "controller",
			call: []string{"con"},
			desc: " - Test controller",
			run: func(s []string) {
				for {com.ReadController()}
			},
		},
	}

	// Show help if no args
	if len(args) == 0 {
		helpFunc(cmds)
	}
	
	if a := findCommand(args[0], cmds); a.name != "" {
		a.run(resolveArgs(args[1:], a.params))
		return
	}
	fmt.Println("Command not found")

	// // Loop through user input arguments
	// for i, a := range args {

	// 	// Check for argument
	// 	if _, e := strconv.Atoi(a); !strings.HasPrefix(a, "-") || e == nil {
	// 		continue
	// 	}

	// 	// Find and execute arguments
	// 	arg := findArg(a, _args)
	// 	if arg.name != "" {
	// 		arg.run(resolveArgs(args[i + 1:], arg.params))
	// 	}
	// }
}
