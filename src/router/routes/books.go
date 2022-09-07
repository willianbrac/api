package routes

import (
	"api/src/controllers"
	"net/http"
)

var BooksRoutes = []Route{
	{
		URI:    "/books",
		Method: http.MethodPost,
		Function: controllers.CreateBook,
		RequireAuthentication: true,
	},
	{
		URI:    "/books",
		Method: http.MethodGet,
		Function: controllers.ListBooks,
		RequireAuthentication: false,
	},
	{
		URI:    "/books/{bookId}",
		Method: http.MethodGet,
		Function: controllers.GetBook,
		RequireAuthentication: true,
	},
	{
		URI:    "/books/{bookId}",
		Method: http.MethodPut,
		Function: controllers.UpdateBook,
		RequireAuthentication: true,
	},
	{
		URI:    "/books/{bookId}",
		Method: http.MethodDelete,
		Function: controllers.DeleteBook,
		RequireAuthentication: true,
	},
}