package models

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"pp/api/utils"
	"strconv"
)

type UserID int64

// Role schema of the roles table
type Role struct {
	ID        UserID  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedAt string `json:"createdAt" db:"created_at"`
	UpdatedAt string `json:"updatedAt" db:"updated_at"`
}

func GetRole(db * sqlx.DB, id int64) (*Role, error) {
	var role Role
	sqlStatement := `SELECT * FROM roles WHERE id=($1)`
	err := db.Get(&role, sqlStatement, strconv.FormatInt(id, 10))

	if err != nil {
		return nil, err
	}

	return &role, err
}

func CreateRole(db * sqlx.DB, name string) (*Role, error) {
	var role Role
	sqlStatement := `INSERT INTO roles (name, created_at, updated_at) VALUES($1, now(), now()) RETURNING id, name, created_at, updated_at`

	if err := db.QueryRowx(sqlStatement, name).Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt); err != nil {
		return nil, err
	}
	return &role, nil
}

func UpdateRole(db * sqlx.DB, id UserID, name string) (*Role, error) {
	var role Role
	sqlStatement := `UPDATE roles SET name=($2), updated_at=now() WHERE id=($1) RETURNING id, name, created_at, updated_at`

	if err := db.QueryRowx(sqlStatement, id, name).Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt); err != nil {
		return nil, err
	}
	return &role, nil
}

func DeleteRole(db * sqlx.DB, id UserID) error {
	sqlStatement := `DELETE FROM roles WHERE id=($1)`
	_, err := db.Exec(sqlStatement, id)
	return err
}

func GetRoles(db * sqlx.DB, limit, offset int) ([]Role, error) {
	sqlStatement := `SELECT id, name, created_at, updated_at FROM roles LIMIT ($1) OFFSET ($2)`
	var roles []Role
	if err := db.Select(&roles, sqlStatement, limit, offset); err != nil {
		return nil, err
	}
	return roles, nil
}

func CheckGetRolesError(error error) error {
	if errors.Is(error, sql.ErrNoRows) {
		return utils.UserNotFoundError
	} else {
		return utils.DBInternalError
	}
}
