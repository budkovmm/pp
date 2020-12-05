package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"

	"pp/api/models"
	"pp/api/utils"
)

var GetRole = func(db* sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}

		gotRole, err := models.GetRole(db, int64(id))
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "User not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, gotRole)
	}
}

var GetRoles = func(db* sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, _ := strconv.Atoi(r.FormValue("count"))
		start, _ := strconv.Atoi(r.FormValue("start"))

		if count > 10 || count < 1 {
			count = 10
		}
		if start < 0 {
			start = 0
		}

		products, err := models.GetRoles(db, start, count)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, products)
	}
}

var CreateRole = func(db* sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var role models.Role
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&role); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		createdRole, err := models.CreateRole(db, role.Name)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, createdRole)
	}
}

var UpdateRole = func(db* sqlx.DB) http.HandlerFunc	{
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}

		var role models.Role
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&role); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		role.ID = int64(id)

		updatedRole, err := models.UpdateRole(db, role.ID, role.Name)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, updatedRole)
	}
}

func DeleteRole(db* sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
			return
		}

		role := models.Role{ID: int64(id)}
		if err := models.DeleteRole(db, role.ID); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	}
}