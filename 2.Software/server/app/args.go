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
	run    func(ctx *CommandCTX)
	params int
}

type CommandCTX struct {
	 args []string
	 waitGroup *sync.WaitGroup
}

func (ctx *CommandCTX) Float64Arg(arg string) float64 {
	f, _ := strconv.ParseFloat(arg, 64)
	return f
}

// Make sure that args will not be empty so the program won't throw ugly errors :)
func resolveArgs(args []string, params int) []string {
	if len(args) >= params {
		return args
	}
	util.PrintExt("Not enought arguments. Use 'help' command for help.")
	return nil
}

// Find and return argument based on input if it exixst, otherwise return empty argument with no name
func findCommand(cmd string, cmds []Command) Command {
	for _, c := range cmds {
		if strings.Compare(c.name, cmd) == 0 || util.StringArrayHas(c.call, cmd) {
			return c
		} 
	}
	return Command{name: ""}
}

// Help menu
func helpFunc(cmds []Command) {
	out := Version + "\nby vizn3r\n\n\nArguments:\n"
	for _, c := range cmds {
		out += "	" + strings.Join(append([]string{c.name}, c.call...), " / ") + " " + c.desc + "\n"
	}
	fmt.Println(out)
}

// Resolves user arguments
func ResolveArgs(args []string) {

	// Application commands
	var cmds = make([]Command, 0)
	cmds = []Command{
		{name: "\nTEST: "},
		{
			name: "test",
			desc: "[...args] - test command",
			call: []string{"t", "tst"},
			run: func(ctx *CommandCTX) {
				fmt.Println(ctx.args)
			},
		},

		{name: "\nMATH: "},
		{
			name: "degtostep",
			desc: "[angles, dest, step] - Convert degrees to steps",
			call: []string{"dts"},
			params: 3,
			run: func(ctx *CommandCTX) {
				s0, _ := strconv.ParseFloat(ctx.args[0], 64)
				s1, _ := strconv.ParseFloat(ctx.args[1], 64)
				s2, _ := strconv.ParseFloat(ctx.args[2], 64)
				fmt.Println(arm.AngleToSteps(s0, s1, s2))
			},
		},
		{
			name: "steptodeg",
			desc: "[angle, steps, step] - Convert steps to degrees",
			call: []string{"std"},
			params: 3,
			run: func(ctx *CommandCTX) {
				s0, _ := strconv.ParseFloat(ctx.args[0], 64)
				s1, _ := strconv.ParseFloat(ctx.args[1], 64)
				s2, _ := strconv.ParseFloat(ctx.args[2], 64)
				fmt.Println(arm.StepsToAngle(s0, s1, s2))
			},
		},

		{name: "\nARM CONTROL: "},
		{
			name: "code",
			desc: "[code, ...args] | [filePath] - Execute arm code",
			call: []string{"c"},
			params: 1,
			run: func(ctx *CommandCTX) {
				arm.ResolveCode(ctx.args)
			},
		},

		{name: "\nAPPLICATION"},
		{
			name: "http",
			desc: "- Start HTTP server",
			run: func(ctx *CommandCTX) {
				c := ConfigFromFile()
				ctx.waitGroup.Add(1)
				go func(){
					defer ctx.waitGroup.Done()
					com.StartHTTP(c.HTTP)
				}()
			},
		},
		{
			name: "btconnect",
			call: []string{"btc"},
			desc: "- Connect to bluetooth device",
			run: func(ctx *CommandCTX) {
				com.ConnectBT()
			},
		},
		{
			name: "listen",
			call: []string{"l"},
			desc: "[baud] - Print serial data from port",
			run: func(ctx *CommandCTX) {
				// for {com.ReadSerial()}
			},
		},
		{
			name: "controller",
			call: []string{"con"},
			desc: " - Test controller",
			run: func(ctx *CommandCTX) {
				// for {com.ReadController()}
			},
		},
	}

	// Show help if no args
	if len(args) == 0 {
		helpFunc(cmds)
		return
	}
	
	// Find command and execute
	if a := findCommand(args[0], cmds); a.name != "" {
		ctx := new(CommandCTX)
		if ctx.args = resolveArgs(args[1:], a.params); ctx.args == nil {
			return
		}

		ctx.waitGroup = new(sync.WaitGroup)
		a.run(ctx)
		ctx.waitGroup.Wait()
		return
	}
	util.PrintExt("Command not found")
}
