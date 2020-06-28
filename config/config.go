package config

import (
	"log"
	"os"
)

// Config is a setting of the App.
type Config struct {
	Port int
	DB   *DBconfig
	Log  *LogConfig
}

// GetConfig returns an App configuration.
func GetConfig() *Config {
	return &Config{
		Port: 8081,
		DB: &DBconfig{
			Host:     "localhost",
			DBname:   "project_00",
			User:     "user_00",
			Password: "159258",
			Port:     5432,
		},
		Log: &LogConfig{
			Logger: log.New(os.Stdout, "", log.LstdFlags),
		},
	}
}
