package app

import "firmware/util"

var MAINCONFIGPATH = "./config.json"

// Configuration for this application
type Config struct {
	// Configurations for other parts of program
	ServerBin   string
	
	// Configurations from other programs
	FirmwareBin string
}

func ConfigToFile(data *Config) {
	util.ToJSON(data, MAINCONFIGPATH)
}

func ConfigFromFile() *Config {
	c := new(Config)
	err := util.ParseJSON(&c, MAINCONFIGPATH)
	if err != nil {
		ConfigToFile(c)
	}
	return c
}