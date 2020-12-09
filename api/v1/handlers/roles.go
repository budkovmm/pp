package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
	"net/http"
	"strconv"

	"pp/api/models"
	"pp/api/utils"
)

type RoleCreateRequest struct {
	Name string `validate:"nonzero,nonnil,min=3,max=40,regexp=^[a-zA-Z0-9_.-]*$"`
}

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

func parseCreateRolePayload(r *http.Request) (*RoleCreateRequest, error) {
	rcr := new(RoleCreateRequest)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(rcr); err != nil {
		return nil, utils.InvalidRequestPayload
	}
	defer r.Body.Close()
	return rcr, nil
}

func validateCreateRoleRequest(rcr *RoleCreateRequest) error {
	if err := validator.Validate(rcr); err != nil {
		return utils.GetValidationError(err.Error())
	}
	return nil
}

func GetRole(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDbFromContext(r.Context())

	if errors.Is(utils.NoDbInContext, err) {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

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

func GetRoles(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDbFromContext(r.Context())

	if errors.Is(utils.NoDbInContext, err) {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	limit, offset := getLimitOffset(r)
	checkLimitOffset(&limit, &offset)
	products, err := models.GetRoles(db, limit, offset)
	if err != nil {
		err = utils.DBInternalError
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, products)
}

func CreateRole(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDbFromContext(r.Context())

	if errors.Is(utils.NoDbInContext, err) {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	roleCreationRequest, err  := parseCreateRolePayload(r)
	if errors.Is(err, utils.InvalidRequestPayload) {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateCreateRoleRequest(roleCreationRequest); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	createdRole, err := models.CreateRole(db, roleCreationRequest.Name)
	if err != nil {
		err = utils.DBInternalError
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdRole)
}

func UpdateRole(w http.ResponseWriter, r *http.Request)	{
	db, err := utils.GetDbFromContext(r.Context())

	if errors.Is(utils.NoDbInContext, err) {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	id, err := getIdFromUrl(r)

	if errors.Is(err, utils.InvalidUserIdError) {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	roleCreationRequest, err := parseCreateRolePayload(r)

	if errors.Is(err, utils.InvalidRequestPayload) {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	updatedRole, err := models.UpdateRole(db, id, roleCreationRequest.Name)

	if err != nil {
		err = utils.DBInternalError
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, updatedRole)
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDbFromContext(r.Context())

	if errors.Is(utils.NoDbInContext, err) {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	id, err := getIdFromUrl(r)
	if errors.Is(err, utils.InvalidUserIdError) {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := models.DeleteRole(db, id); err != nil {
		err = utils.DBInternalError
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}