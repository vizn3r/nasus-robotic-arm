package com

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type HTTPConf struct {
	Port string
}

func StartHTTP(conf HTTPConf) {
	fmt.Println("Port: " + conf.Port)
	server := fiber.New()
	server.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello world!"))
	})
	server.Listen(":" + conf.Port)
}