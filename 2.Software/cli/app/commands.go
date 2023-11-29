package app

import (
	"cli/com"
	"cli/ext"
	"fmt"
	"os/exec"
)

var CMDS = make([]Command, 0)

func Cmds() { CMDS = []Command{
	{
		name: "help",
		desc: "- Prints this help menu",
		call: []string{"h"},
		run: func(ctx *CommandCTX) {
			helpFunc(CMDS)
		},
	},

	{name: "\nAPPLICATION:"},
	{
		name: "config",
		desc: "[serverPath] - Generate config in current directory",
		call: []string{"conf"},
		run: func(ctx *CommandCTX) {
			c := ext.ConfigFromFile()
			if a := ctx.args; len(a) != 0 {
				c.ServerBin = a[0]
			}
			c.ServerBin = "./server.exe"
			ext.ConfigToFile(c)
			fmt.Println("New config created at", ext.MAINCONFIGPATH)
		},
		params: 0,
	},
	{
		name: "http",
		desc: "- Start HTTP server",
		call: []string{},
		run: func(*CommandCTX) {
			c := ext.ConfigFromFile()
			b, _ := exec.Command(c.ServerBin).Output()
			fmt.Println(string(b))
		},
		params: 0,
	},
	{
		name: "send",
		desc: "[message, port] - Send data to firmware",
		call: []string{"s"},
		run: func(c *CommandCTX) {
			if len(c.args) == 1 {
				c.args = append(c.args, ":8080")
			}
			com.Send(c.args[0], c.args[1])
		},
		params: 0,
	},
	// {
	// 	name: "btconnect",
	// 	call: []string{"btc"},
	// 	desc: "- Connect to bluetooth device",
	// 	run: func(*CommandCTX) {
	// 		com.ConnectBT()
	// 	},
	// },
	// {
	// 	name: "listen",
	// 	call: []string{"l"},
	// 	desc: "[baud] - Print serial data from port",
	// 	run: func(ctx *CommandCTX) {
	// 		for {com.ReadSerial()}
	// 	},
	// },
	// {
	// 	name: "controller",
	// 	call: []string{"con"},
	// 	desc: " - Test controller",
	// 	run: func(ctx *CommandCTX) {
	// 		for {com.ReadController()}
	// 	},
	// },
}}