package app

import (
	"fmt"
	"server/arm"
	"server/util"
	"strconv"
	"strings"
)

var Version = "Nasus Server v0.0.1" // VERSION OF SOFTWARE

type arg struct {
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
func findArg(a string, args []arg) arg {
	for _, _a := range args {
		if strings.Compare(_a.name, a[1:]) == 0 || util.StringArrayHas(_a.call, a[1:]) {
			return _a
		} 
	}
	fmt.Println("Argument not found.")
	helpFunc(args)
	return arg{name: ""}
}

// Help menu
func helpFunc(args []arg) {
	out := Version + "\nby vizn3r\n\nArguments:\n"
	for _, a := range args {
		out += "	" + a.name + " " + a.desc + "\n"
	}
	fmt.Println(out)
}

// Resolves user arguments
func ResolveArgs(args []string) {

	// Application arguments
	var _args = make([]arg, 0)
	_args = []arg{
		{
			name: "test",
			desc: "[...args] - test command",
			call: []string{"t", "tst"},
			run: func(s []string) {
				fmt.Println(s)
			},
		},

		// Math arguments
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

		// Arm control
		{
			name: "code",
			desc: "[code, ...args] - Execute arm code",
			call: []string{"c"},
			params: 1,
			run: func(s []string) {
				arm.ResolveCode(s[0], s[1:])
			},
		},

		// Config
		// {
		// 	name: "config",
		// 	desc: "[option, value] - Configure a value",
		// 	call: []string{"conf", "cf", "set"},
		// 	params: 2,
		// 	run: func(s []string) {
		// 		c := ConfigFromFile()

		// 		c.UpdateFile(*c)
		// 	},
		// },
	}

	// Show help if no args
	if len(args) == 0 {
		helpFunc(_args)
	}
	
	// Loop through user input arguments
	for i, a := range args {

		// Check for argument
		if _, e := strconv.Atoi(a); !strings.HasPrefix(a, "-") || e == nil {
			continue
		}

		// Find and execute arguments
		arg := findArg(a, _args)
		if arg.name != "" {
			arg.run(resolveArgs(args[i + 1:], arg.params))
		}
	}
}
