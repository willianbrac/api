package routes

import (
	"api/src/router/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: controllers.ListUsers,
		RequireAuthentication: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodGet,
		Function: controllers.GetUser,
		RequireAuthentication: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		RequireAuthentication: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		RequireAuthentication: false,
	},
}