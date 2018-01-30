package users

import (
	"net/http"

	"github.com/geektimus/parking-mgmt/parking-mgmt-be/logger"
	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/users",
		controller.Index,
	},
	Route{
		"GetUser",
		"GET",
		"/api/users/{id}",
		controller.GetUser,
	},
	Route{
		"AddUser",
		"POST",
		"/api/users",
		controller.AddUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/api/users",
		controller.UpdateUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/api/users/{id}",
		controller.DeleteUser,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
