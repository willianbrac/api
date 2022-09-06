package routes

import (
	"api/src/controllers"
	"net/http"
)

var LoginRoute = Route{
	URI:                   "/signIn",
	Method:                http.MethodPost,
	Function:              controllers.Login,
	RequireAuthentication: false,
}