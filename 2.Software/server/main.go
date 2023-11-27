package main

import (
	"fmt"
	"os"
	"os/exec"
	"server/auth"
	"server/config"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Should redo this to support arguments
	// Maybe I should make an argument library, which I have so I don't need to
	firmwarePath := ""
	if c, args := config.ConfigFromFile(), os.Args; c.FirmwareBin != "" {
		firmwarePath = c.FirmwareBin
	} else if len(args) == 2 && strings.ToLower(args[1]) == "conf" {
		c.FirmwareBin = "./firmware.exe"
		config.ConfigToFile(c)
		fmt.Println("New config created at", config.MAINCONFIGPATH)
		os.Exit(0)
	} else if len(args) == 2 {
		c.FirmwareBin = args[1]	
		config.ConfigToFile(c)
		fmt.Println("New config created at", config.MAINCONFIGPATH)
		os.Exit(0)
	} else {
		fmt.Println("No firmware binary/executable. Provide a [path] or enter 'conf' to auto-generate config.")
		os.Exit(0)
	}
	

	app := fiber.New()
	app.Post("/*", func(c *fiber.Ctx) error {
		if !auth.Auth(c.GetReqHeaders()["Authorization"][0]) {
			return c.SendStatus(401)
		}
		
		cli := exec.Command(firmwarePath)
		cli.Args = strings.Split(c.Path(), "/")
		out, err := cli.Output()
		if err != nil {
			return c.SendString(fmt.Sprintf("%v", err) + "\n\n👆 That probably means that I fucked up somewhere in code :D")
		}
		return c.Send(out)
	})
	app.Listen(":8080")
}