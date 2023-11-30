package arm

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Mainly because I'm lazy to manually write documentation, so this is for doc autogen.
type Code struct {
	Desc string // The whole description with everything
	Params []string
	Run func(args ...string)
}

var GCodes = []Code { // For everything related with movement etc.
	{
		Desc: "test",
		Run: func(args ...string) {
			fmt.Println(args)
		},
	},
}

var MCodes = []Code { // For everyting else
	// {
	// 	Desc: "Start TCP Server",
	// 	Params: []string{"[port] - Server port"},
	// 	Run: func(args ...string) {
	// 		com.CLIServer.IsEnabled = true
	// 	},
	// },
}

var Version = "Nasus Firmware v0.0.1" // VERSION OF SOFTWARE

// Resolves user arguments
func ResolveArgs(args []string) []string {
	// Show help if no args
	if len(args) == 0 {
		fmt.Println("Invalid.")
		return nil
	}
	return args
}

func StringToInt(s string) int {
	if f, e := strconv.Atoi(s); e == nil {
		return f
	}
	return -1
}

func ExecCode(code []string) string {
	if len(code[0]) != 2 {
		return "Invalid."
	}
	if strings.ToLower(code[0][0:1]) == "g" {
		i := StringToInt(code[0][1:2])
		if i == -1 || i >= len(GCodes) {
			return "Invalid."
		}
		GCodes[i].Run(code[1:]...)
	} else if strings.ToLower(code[0][0:1]) == "m" {
		i := StringToInt(code[0][1:2])
		if i == -1 || i >= len(MCodes) {
			return "Invalid."
		}
		MCodes[i].Run(code[1:]...)
	}
	return "Ok."
}

func DocGen() {
	var data string
	header := "# " + Version + " Documentation"
	codes := "## GCode List"
	data += header + "\n\n" + codes + "\n\n"
	for i, g := range GCodes {
		data += "### G" + strconv.Itoa(i) + "\n\n"
		data += "> " + g.Desc + "\n"
	}
	for i, m := range MCodes {
		data += "### M" + strconv.Itoa(i) + "\n\n"
		data += "> " + m.Desc + "\n"
	}

	err := os.WriteFile("./FIRMWARE.md", []byte(data), 0777)
	if err != nil {
		return
	}
}