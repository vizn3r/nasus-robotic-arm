package app

import (
	"firmware/arm"
	"firmware/com"
	"firmware/util"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
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
	 CMDS []Command
}

func (ctx *CommandCTX) Float64Arg(argIndex int) float64 {
	f, _ := strconv.ParseFloat(ctx.args[argIndex], 64)
	return f
}

// Make sure that args will not be empty so the program won't throw ugly errors :)
func parseArgs(args []string, params int) []string {
	if len(args) >= params {
		return args
	}
	fmt.Println("Not enought arguments. Use 'help' command for help.")
	return nil
}

// Find and return argument based on input if it exixst, otherwise return empty argument with no name
func FindCommand(cmd string) Command {
	for _, c := range CMDS {
		if strings.Compare(c.name, cmd) == 0 || util.StringArrayHas(c.call, cmd) {
			return c
		} 
	}
	return Command{name: ""}
}

// Help menu
func helpFunc(cmds []Command) {
	out := Version + "\nby vizn3r\n\n\nCommands:\n"
	for _, c := range cmds {
		out += "	" + strings.Join(append([]string{c.name}, c.call...), " / ") + " " + c.desc + "\n"
	}
	fmt.Println(out)
}

var CMDS = make([]Command, 0)

// Resolves user arguments
func ResolveArgs(args []string) {
	
	// Application commands
	CMDS = []Command{
		{
			name: "help",
			desc: "- Prints this help menu",
			call: []string{"h"},
			run: func(ctx *CommandCTX) {
				helpFunc(CMDS)
			},
		},

		{name: "\nMATH: "},
		{
			name: "degtostep",
			desc: "[angles, dest, step] - Convert degrees to steps",
			call: []string{"dts"},
			params: 3,
			run: func(ctx *CommandCTX) {
				fmt.Println(arm.AngleToSteps(ctx.Float64Arg(0), ctx.Float64Arg(1), ctx.Float64Arg(2)))
			},
		},
		{
			name: "steptodeg",
			desc: "[angle, steps, step] - Convert steps to degrees",
			call: []string{"std"},
			params: 3,
			run: func(ctx *CommandCTX) {
				fmt.Println(arm.StepsToAngle(ctx.Float64Arg(0), ctx.Float64Arg(1), ctx.Float64Arg(2)))
			},
		},

		{name: "\nARM CONTROL: "},
		{
			name: "code",
			desc: "[{code, ...args} || {filePath}] - Execute code or read .gcode file",
			call: []string{"c"},
			params: 1,
			run: func(ctx *CommandCTX) {
				arm.ResolveCode(ctx.args)
			},
		},

		{name: "\nAPPLICATION:"},
		{
			name: "config",
			desc: "[serverPath] - Generate config in current directory",
			call: []string{"conf"},
			run: func(ctx *CommandCTX) {
				c := ConfigFromFile()
				if a := ctx.args; len(a) != 0 {
					c.ServerBin = a[0]
				}
				c.ServerBin = "./server.exe"
				ConfigToFile(c)
				fmt.Println("New config created at", MAINCONFIGPATH)
			},
			params: 0,
		},
		{
			name: "http",
			desc: "- Start HTTP server",
			call: []string{},
			run: func(*CommandCTX) {
				c := ConfigFromFile()
				b, _ := exec.Command(c.ServerBin).Output()
				fmt.Println(string(b))
			},
			params: 0,
		},
		{
			name: "btconnect",
			call: []string{"btc"},
			desc: "- Connect to bluetooth device",
			run: func(*CommandCTX) {
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
	if len(args) == 0 || args[0] == "" {
		helpFunc(CMDS)
		return
	}
	
	// Find command and execute
	if a := FindCommand(args[0]); a.name != "" {
		ctx := new(CommandCTX)
		if ctx.args = parseArgs(args[1:], a.params); ctx.args == nil {
			return
		}

		ctx.CMDS = CMDS
		a.run(ctx)
		return
	}
	fmt.Println("Command '" + args[0] + "' not found")
}
