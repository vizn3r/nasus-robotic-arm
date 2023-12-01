package com

import (
	"firmware/arm"
	"fmt"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type CLIConf struct {
	Port string
}

type CLI struct {
	COM
	Conf CLIConf
}

var CLIServer = new(CLI)

// Comunicating with CLI
func (cli *CLI) StartCLI(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	if !cli.Start("CLI Server") {
		return 
	}
	defer fmt.Println("'CLI Server' stopped.")
	cli.isRunning = true
	
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.Send([]byte(string(arm.ExecCode(arm.ResolveArgs(strings.Split(c.Path()[1:], "/"))))))
	})

	app.Listen(cli.Conf.Port)
}