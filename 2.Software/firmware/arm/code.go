package arm

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Version = "Nasus Firmware v0.0.1" // VERSION OF SOFTWARE

// Mainly because I'm lazy to manually write documentation, so this is for doc autogen.
type Code struct {
	Desc string // The whole description with everything
	Params []string
	Run func(args ...string)
}

var GCodes = []Code { // For everything related with movement etc.
	
}

var MCodes = []Code { // For everyting else
	
}

var TCodes = []Code {
	{
		Desc: "For testing functions",
		Run: func(args ...string) {
			m := new(Motor)
			m.StepAngle = 1.8
			m.STEP = 17
			m.RotateDeg(90, 1, 1)
		},
	},
}

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
	} else if strings.ToLower(code[0][0:1]) == "t" {
		i := StringToInt(code[0][1:2])
		if i == -1 || i >= len(TCodes) {
			return "Invalid."
		}
		TCodes[i].Run(code[1:]...)
	}
	return "Ok."
}

func DocGen() {
	var data string
	header := "# " + Version + " Documentation"
	codes := "## GCode List"
	data += header + "\n\n" + codes + "\n\n"
	for i, g := range GCodes {
		data += "### G - Motion and function" + strconv.Itoa(i) + "\n\n"
		data += "> **Description**" + "\n> \n" + "> " + g.Desc + "\n"
	}
	for i, m := range MCodes {
		data += "### M - Operations not involving movements" + strconv.Itoa(i) + "\n\n"
		data += "> " + m.Desc + "\n"
	}
	err := os.WriteFile("./FIRMWARE.md", []byte(data), 0777)
	if err != nil {
		return
	}
}