package com

import (
	"github.com/gofiber/fiber/v2"
)

type HTTPConf struct {
	Port string
}

func StartHTTP(conf HTTPConf) {
	server := fiber.New()
	server.Listen(":" + conf.Port)
}