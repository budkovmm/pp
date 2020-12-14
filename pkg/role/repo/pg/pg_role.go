package pg

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"pp/pkg/domain"
	"pp/pkg/model"
)

type pgRoleRepository struct {
	DbConn *sqlx.DB
}

func NewPGRoleRepository(DbConn *sqlx.DB) domain.RoleRepository {
	return &pgRoleRepository{DbConn}
}

func (rr *pgRoleRepository) GetById(id model.RoleID) (res *model.Role, err error) {
	var role model.Role
	sqlStatement := `SELECT * FROM roles WHERE id=($1)`
	err = rr.DbConn.Get(&role, sqlStatement, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.RoleNotFoundError
		} else {
			return nil, domain.DBInternalError
		}
	}

	return &role, err
}

func (rr *pgRoleRepository) Create(name model.RoleName) (*model.Role, error) {
	var role model.Role
	sqlStatement := `INSERT INTO roles (name, created_at, updated_at) VALUES($1, now(), now()) RETURNING id, name, created_at, updated_at`

	if err := rr.DbConn.QueryRowx(sqlStatement, name).Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt); err != nil {
		return nil, err
	}
	return &role, nil
}

func (rr *pgRoleRepository) Update(id model.RoleID, name model.RoleName) (*model.Role, error) {
	var role model.Role
	sqlStatement := `UPDATE roles SET name=($2), updated_at=now() WHERE id=($1) RETURNING id, name, created_at, updated_at`

	err := rr.DbConn.QueryRowx(sqlStatement, id, name).Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.RoleNotFoundError
		} else {
			return nil, domain.DBInternalError
		}
	}
	return &role, nil
}

func (rr *pgRoleRepository) Delete(id model.RoleID) error {
	sqlStatement := `DELETE FROM roles WHERE id=($1)`
	_, err := rr.DbConn.Exec(sqlStatement, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.RoleNotFoundError
		} else {
			return domain.DBInternalError
		}
	}
	return nil
}

func (rr *pgRoleRepository) GetAll(limit, offset int) ([]model.Role, error) {
	sqlStatement := `SELECT id, name, created_at, updated_at FROM roles ORDER BY created_at LIMIT ($1) OFFSET ($2)`
	var roles []model.Role
	if err := rr.DbConn.Select(&roles, sqlStatement, limit, offset); err != nil {
		return nil, err
	}
	return roles, nil
}
