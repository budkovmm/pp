package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"

	"pp/api/models"
	"pp/api/utils"
)

func getIdFromUrl(r *http.Request) (models.UserID, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return -1, utils.InvalidUserIdError
	}
	return models.UserID(int64(id)), nil
}

func getLimitOffset(r *http.Request) (limit, offset int){
	limit, _ = strconv.Atoi(r.FormValue("limit"))
	offset, _ = strconv.Atoi(r.FormValue("offset"))
	return limit, offset
}

func checkLimitOffset(limit, offset *int) {
	if *limit > 10 || *limit < 1 {
		*limit = 10
	}
	if *offset < 0 {
		*offset = 0
	}
}

func parseCreateRolePayload(r *http.Request, role *models.Role) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(role); err != nil {
		return utils.InvalidRequestPayload
	}
	defer r.Body.Close()
	return nil
}

var GetRole = func(db* sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getIdFromUrl(r)
		if errors.Is(err, utils.InvalidUserIdError) {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		gotRole, err := models.GetRole(db, int64(id))

		if err != nil {
			err = models.CheckGetRolesError(err)
			if errors.Is(err, utils.UserNotFoundError) {
				utils.RespondWithError(w, http.StatusNotFound, err.Error())
			}
			if errors.Is(err, utils.DBInternalError) {
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, gotRole)
	}
}

var GetRoles = func(db* sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, offset := getLimitOffset(r)
		log.Printf("Limit is %d | Offset is %d", limit, offset)
		checkLimitOffset(&limit, &offset)
		log.Printf("Limit is %d | Offset is %d", limit, offset)
		products, err := models.GetRoles(db, limit, offset)
		if err != nil {
			err = utils.DBInternalError
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, products)
	}
}

var CreateRole = func(db* sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var role models.Role
		if err := parseCreateRolePayload(r, &role); errors.Is(err, utils.InvalidRequestPayload) {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		}

		createdRole, err := models.CreateRole(db, role.Name)
		if err != nil {
			err = utils.DBInternalError
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, createdRole)
	}
}

var UpdateRole = func(db* sqlx.DB) http.HandlerFunc	{
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getIdFromUrl(r)
		if errors.Is(err, utils.InvalidUserIdError) {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		var role models.Role
		if err := parseCreateRolePayload(r, &role); errors.Is(err, utils.InvalidRequestPayload) {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		role.ID = id

		updatedRole, err := models.UpdateRole(db, role.ID, role.Name)
		if err != nil {
			err = utils.DBInternalError
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, updatedRole)
	}
}

func DeleteRole(db* sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getIdFromUrl(r)
		if errors.Is(err, utils.InvalidUserIdError) {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		role := models.Role{ID: id}
		if err := models.DeleteRole(db, role.ID); err != nil {
			err = utils.DBInternalError
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	}
}