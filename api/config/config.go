package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	ConnectionStringDB    = ""   // It's the connection string for MySQL.
	ConnectionPort        = 0    // It's the port where the API will run.
	ConnectionPortDefault = 9000 // It's a default port in case of error reading the .env file.
)

// LoadConfig load environment variable settings inside the api.
func LoadConfig() {

	var error error

	if error = godotenv.Load("../../.env"); error != nil { // Tries to read the .env file retrieving the environment variables.
		log.Fatal(error) // If error stop the application.
	}

	ConnectionPort, error = strconv.Atoi(os.Getenv("PORT_API"))
	if error != nil {
		ConnectionPort = ConnectionPortDefault
	}

	ConnectionStringDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", // Using sprintf to format a connection string.
		os.Getenv("USER_DB"),
		os.Getenv("PASSWORD_DB"),
		os.Getenv("NAME_DB"),
	)
}
