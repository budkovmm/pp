package tests
//
//var ErrUserNotFound = errors.New("user nt found")
//
//func GetRoleByID(id int64) (Role, error) {
//	gotRole, err := models.GetRole(r.Context(), int64(id))
//	if err != nil {
//		switch err {
//		case sql.ErrNoRows:
//			return nil, ErrUserNotFound
//		default:
//			return nil, errors.Wrap("SQL Exception", err)
//		}
//		return
//	}
//	return
//}
//
//
//package handlers
//func GetRole(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
//		return
//	}
//
//	if err = GetRoleByID(id int64) ;err != nil {
//		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
//		return
//	}
//	utils.RespondWithJSON(w, http.StatusOK, gotRole)
//}
//
//
//
//
//
//
//
//
//
//
//package roles
//const (
//	maxLimit = 10
//)
//
//func checkBorders(l, s int) (int, int) {
//	if count > maxLimit || count < 1 {
//		count = maxLimit
//	}
//	if start < 0 {
//		start = 0
//	}
//}
//
//func GetRoles(limit, offset int) ([]Roles, error) {
//
//	l, o := checkBorders()
//	products, err := models.GetRoles(r.Context(), start, count)
//	if err != nil {
//		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//}
//
//
//package handlers
//
//func GetRoles(w http.ResponseWriter, r *http.Request) {
//	count, _ := strconv.Atoi(r.FormValue("count"))
//	start, _ := strconv.Atoi(r.FormValue("start"))
//
//	roles.GetAll(count, start)
//
//
//	utils.RespondWithJSON(w, http.StatusOK, products)
//}