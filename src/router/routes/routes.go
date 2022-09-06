package routes

import (
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
	for _, route := range routes{
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}	
	return r
}