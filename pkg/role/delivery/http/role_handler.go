package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"pp/pkg/domain"
	"pp/pkg/model"
	"pp/pkg/utils"
	"strconv"
)

type RoleRequest struct {
	Name model.RoleName
}

func parseRolePayload(r *http.Request) (*RoleRequest, error) {
	rp := new(RoleRequest)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(rp); err != nil {
		return nil, domain.InvalidRequestPayload
	}
	defer r.Body.Close()
	return rp, nil
}

func getLimitOffset(r *http.Request) (limit, offset int){
	limit, _ = strconv.Atoi(r.FormValue("limit"))
	offset, _ = strconv.Atoi(r.FormValue("offset"))
	return limit, offset
}

func getIdFromUrl(r *http.Request) (model.RoleID, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return -1, domain.InvalidUserIdError
	}
	return model.RoleID(int64(id)), nil
}

type RoleHandler struct {
	RoleUseCase domain.RoleUseCase
}

func NewRoleHandler(r *mux.Router, us domain.RoleUseCase) {
	handler := &RoleHandler{
		RoleUseCase: us,
	}
	r.HandleFunc("/role/{id:[0-9]+}", handler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/role/{id:[0-9]+}", handler.Update).Methods(http.MethodPut)
	r.HandleFunc("/role/{id:[0-9]+}", handler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/role", handler.Create).Methods(http.MethodPost)
	r.HandleFunc("/role", handler.GetAll).Methods(http.MethodGet)
}

func (rh *RoleHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromUrl(r)

	if errors.Is(err, domain.InvalidUserIdError) {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	gotRole, err := rh.RoleUseCase.GetById(id)

	if err != nil {
		if errors.Is(err, domain.RoleNotFoundError) {
			utils.RespondWithError(w, http.StatusNotFound, err.Error())
		}
		if errors.Is(err, domain.DBInternalError) {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, gotRole)
}

func (rh *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromUrl(r)

	if errors.Is(err, domain.InvalidUserIdError) {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	roleRequest, err := parseRolePayload(r)

	if errors.Is(domain.InvalidRequestPayload, err) {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedRole, err := rh.RoleUseCase.Update(id, roleRequest.Name)

	if err != nil {
		if errors.Is(err, domain.RoleNotFoundError) {
			utils.RespondWithError(w, http.StatusNotFound, err.Error())
		}
		if errors.Is(err, domain.DBInternalError) {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, updatedRole)
}

func (rh *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromUrl(r)

	if errors.Is(err, domain.InvalidUserIdError) {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = rh.RoleUseCase.Delete(id)

	if err != nil {
		if errors.Is(err, domain.RoleNotFoundError) {
			utils.RespondWithError(w, http.StatusNotFound, err.Error())
		}
		if errors.Is(err, domain.DBInternalError) {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	resultMessage := fmt.Sprintf("Role %d was successfully deleted", id)
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": resultMessage})
}

func (rh *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	roleRequest, err := parseRolePayload(r)

	if errors.Is(domain.InvalidRequestPayload, err) {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	createdRole, err := rh.RoleUseCase.Create(roleRequest.Name)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdRole)
}

func (rh *RoleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	limit, offset := getLimitOffset(r)
	gotRoles, err := rh.RoleUseCase.GetAll(limit, offset)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, gotRoles)

}

