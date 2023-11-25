package app

import (
	"fmt"
	"server/com"
	"server/util"
)

var MAINCONFIGPATH = "./.config/config.json"

// Configuration for this application
type Config struct {
	Test string
	// Configurations for other parts of program
	HTTP com.HTTPConf
}

func ConfigToFile(data* Config) {
	util.ToJSON(*data, MAINCONFIGPATH)
}

func ConfigFromFile() *Config {
	c := new(Config)
	err := util.ParseJSON(c, MAINCONFIGPATH)
	if err != nil {
		fmt.Println("Config created.")
	}
	return c
}

func (c *Config) UpdateFile(conf Config) {
	err := util.ToJSON(conf, MAINCONFIGPATH)
	if err != nil {
		fmt.Println("Config does not exist.")
	}
}