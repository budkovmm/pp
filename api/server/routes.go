package server

import (
	"net/http"
	"pp/api/v1/handlers"
)

const (
	roleByIdPattern = "/role/{id:[0-9]+}"
	rolePattern     = "/role"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}

// Routes - A collection of Routes
type Routes []Route

// APIRoutes - Routes for PP API
var apiRoutes = Routes{
	Route{
		http.MethodGet,
		roleByIdPattern,
		handlers.GetRole,
		false,
	},
	Route{
		http.MethodGet,
		rolePattern,
		handlers.GetRoles,
		false,
	},
	Route{
		http.MethodPost,
		rolePattern,
		handlers.CreateRole,
		false,
	},
	Route{
		http.MethodPut,
		roleByIdPattern,
		handlers.UpdateRole,
		false,
	},
	Route{
		http.MethodDelete,
		roleByIdPattern,
		handlers.DeleteRole,
		false,
	},
}