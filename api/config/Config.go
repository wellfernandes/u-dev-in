package config

import (
	"os"
	"strconv"
	"sync"
)

type Config struct {
	Env      string
	DbConn   string
	RestPort int
}

var config *Config
var doOnce sync.Once

func GetConfig() *Config {

	doOnce.Do(func() {
		config = &Config{
			Env:    os.Getenv("Env"),
			DbConn: os.Getenv("DbConn"),
		}

		restPort := os.Getenv("RestPort")

		intRestPort, err := strconv.Atoi(restPort)
		if err != nil {
			panic(err)
		}

		config.RestPort = intRestPort
	})

	return config
}
