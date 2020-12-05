package server

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewServer(db *sqlx.DB, prefix string) (*mux.Router, error) {
	r := mux.NewRouter()
	apiV1 := r.PathPrefix(prefix).Subrouter()
	for _, v:= range GetRoutes(db) {
		apiV1.Handle(v.Pattern, v.HandlerFunc(db)).Methods(v.Method)
	}
	return apiV1, nil
}