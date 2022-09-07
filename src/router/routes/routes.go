package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := routesUsers
	routes = append(routes, LoginRoute)
	routes = append(routes, BooksRoutes...)

	for _, route := range routes{
		if route.RequireAuthentication {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}	
	return r
}