package app

import (
	"cli/util"
	"fmt"
	"strconv"
	"strings"
)

var Version = "Nasus CLI v0.0.1" // VERSION OF SOFTWARE

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

// Resolves user arguments
func ResolveArgs(args []string) {
	Cmds()

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
