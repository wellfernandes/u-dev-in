package router

import (
	"api/rest/routes"
	"github.com/gorilla/mux"
)

// GenerateRouter returns a router with configured routes
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()          //created the router
	return routes.Config_route(r) //returns the router created and configured
}
