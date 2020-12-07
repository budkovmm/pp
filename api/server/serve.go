package server

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"pp/api/utils"
)

func NewServer(prefix string, db *sqlx.DB) (*mux.Router, error) {
	r := mux.NewRouter()
	apiV1 := r.PathPrefix(prefix).Subrouter()
	for _, v:= range apiRoutes {
		apiV1.Handle(v.Pattern, utils.InjectDB(db, v.HandlerFunc)).Methods(v.Method)
	}
	return apiV1, nil
}