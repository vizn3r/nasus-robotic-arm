package app

import (
	"server/com"
	"server/util"
)

var MAINCONFIGPATH = "./.config.json"

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
		ConfigToFile(c)
	}
	return c
}