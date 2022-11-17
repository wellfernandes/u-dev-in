package router

import "github.com/gorilla/mux"

// GenerateRouter returns a router with configured routes
func GenerateRouter() *mux.Router {
	return mux.NewRouter()
}
