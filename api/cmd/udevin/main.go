package main

import (
	"api/config"
	"api/rest/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Configure the connection to the database.
	config.LoadConfig()

	// Create the router with the configured routes.
	r := router.GenerateRouter()

	fmt.Println("API is running!")

	// Print the connection port to the terminal.
	fmt.Println(fmt.Sprintf("Listening to port %d...", config.ConnectionPort))
	// Starts the http server using the port from the config package.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ConnectionPort), r))
}
