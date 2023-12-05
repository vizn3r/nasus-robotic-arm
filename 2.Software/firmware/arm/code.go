package arm

import (
	"firmware/gpio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Version = "Nasus Firmware v0.0.1" // VERSION OF SOFTWARE

// Mainly because I'm lazy to manually write documentation, so this is for doc autogen.
type Code struct {
	Name string
	Desc string // The whole description with everything
	Params []string
	Run func(args ...string)
}

var GCodes = []Code { // For everything related with movement etc.

}

var MCodes = []Code { // For everyting else
	
}

var Codes = []Code {
	{
		Desc: "For testing functions",
		Run: func(args ...string) {
			fmt.Println("Testing motor:")
			m := new(Motor)
			m.StepAngle = stof(args[0])
			m.STEP = 17
			m.RotateDeg(stof(args[1]), 0, stoi(args[2]))
		},
	},
	{
		Run: func(args ...string) {
			d, err := gpio.Read(17)
			if err != nil {
				fmt.Println("Cannot open")
			}
			fmt.Println(d)	
		},
	},
	{
		Run: func(args ...string) {
			err := gpio.Write(17, args[0])
			if err != nil {
				fmt.Println("Error writing")
				return 
			}
		},
	},
}

func stof(s string) float64 {
	i, e := strconv.ParseFloat(s, 64)
	if e != nil {
		return 0
	}
	return i
}

func stoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		return 0
	}
	return i
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
	for _, c := range Codes {
		if strings.ToLower(code[0]) != strings.ToLower(c.Name) {
			return "Invalid."
		}
		c.Run(code[1:]...)
	}
	return "Ok."
}

func DocGen() {
	var data string
	header := "# " + Version + " Documentation"
	codes := "## GCode List"
	data += header + "\n\n" + codes + "\n\n"
	for _, c := range Codes {
		data += "###" + c.Name + "\n\n"
		data += "> **Description**" + "\n> \n" + "> " + c.Desc + "\n"
	}
	err := os.WriteFile("./FIRMWARE.md", []byte(data), 0777)
	if err != nil {
		return
	}
}