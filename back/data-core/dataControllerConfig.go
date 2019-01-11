package datacore

import (
	"github.com/BurntSushi/toml"
)

const configFileName = "./config.toml"

// DatabaseConfig abase used in toml config file to read the database connection
type DatabaseConfig struct {
	Server     string
	Port       int
	Database   string
	Collection string
}

// Read and parse the configuration file
func (db *DatabaseConfig) Read() error {
	_, err := toml.DecodeFile(configFileName, &db)

	return err
}
