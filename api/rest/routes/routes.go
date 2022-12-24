package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Route represents all API routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// ConfigRoute put all routes inside the router.
func ConfigRoute(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
