package server

import (
	"github.com/jmoiron/sqlx"
	"net/http"
	"pp/api/v1/handlers"
)

const (
	roleByIdPattern = "/role/{id:[0-9]+}"
	rolePattern     = "/role"
)

type HandlerWithDB func(*sqlx.DB) http.HandlerFunc

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc HandlerWithDB
	Protected   bool
	DB          *sqlx.DB
}

// Routes - A collection of Routes
type Routes []Route

// APIRoutes - Routes for PP API
func GetRoutes(db *sqlx.DB) []Route {
	var apiRoutes = Routes{
		Route{
			http.MethodGet,
			roleByIdPattern,
			handlers.GetRole,
			false,
			db,
		},
		Route{
			http.MethodGet,
			rolePattern,
			handlers.GetRoles,
			false,
			db,
		},
		Route{
			http.MethodPost,
			rolePattern,
			handlers.CreateRole,
			false,
			db,
		},
		Route{
			http.MethodPut,
			roleByIdPattern,
			handlers.UpdateRole,
			false,
			db,
		},
		Route{
			http.MethodDelete,
			roleByIdPattern,
			handlers.DeleteRole,
			false,
			db,
		},
	}
	return apiRoutes
}