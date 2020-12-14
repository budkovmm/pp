package domain

import (
	"pp/pkg/model"
)

type RoleRepository interface {
	GetById(id model.RoleID) (res *model.Role, err error)
	Create(name model.RoleName) (res *model.Role, err error)
	Update(id model.RoleID, name model.RoleName) (res *model.Role, err error)
	Delete(id model.RoleID) error
	GetAll(limit, offset int) ([]model.Role, error)
}

type RoleUseCase interface {
	GetById(id model.RoleID) (res *model.Role, err error)
	Create(name model.RoleName) (res *model.Role, err error)
	Update(id model.RoleID, name model.RoleName) (res *model.Role, err error)
	Delete(id model.RoleID) error
	GetAll(limit, offset int) ([]model.Role, error)
}