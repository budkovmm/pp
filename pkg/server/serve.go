package server

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_roleHttpDelivery "pp/pkg/role/delivery/http"
	_roleUseCase "pp/pkg/role/usecase"
	_roleRepo "pp/pkg/role/repo/pg"

)

func NewServer(prefix string, db *sqlx.DB) (*mux.Router, error) {
	r := mux.NewRouter()
	apiV1 := r.PathPrefix(prefix).Subrouter()
	roleRepo := _roleRepo.NewPGRoleRepository(db)
	ru := _roleUseCase.NewRoleUseCase(roleRepo)
	_roleHttpDelivery.NewRoleHandler(apiV1, ru)
	return apiV1, nil
}