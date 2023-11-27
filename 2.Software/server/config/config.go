package config

var MAINCONFIGPATH = "./config.json"

// Configuration for this application
type Config struct {
	// Configurations for other parts of program
	FirmwareBin string

	// Configurations from other programs
	ServerBin string
}

func ConfigToFile(data *Config) {
	ToJSON(data, MAINCONFIGPATH)
}

func ConfigFromFile() *Config {
	c := new(Config)
	err := ParseJSON(&c, MAINCONFIGPATH)
	if err != nil {
		ConfigToFile(c)
	}
	return c
}