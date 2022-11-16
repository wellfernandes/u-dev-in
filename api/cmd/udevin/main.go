package main

import (
	"api/rest/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API is running!")

	r := router.GenerateRouter()

	fmt.Println("Listening to port 5000...")
	log.Fatal(http.ListenAndServe(":5000", r))

}
