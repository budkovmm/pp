package usecase

import (
	"pp/pkg/domain"
	"pp/pkg/model"
)

const (
	DefaultLimitValue = 10
	DefaultOffsetValue = 0
)

func checkLimitOffset(limit, offset *int) {
	if *limit > DefaultLimitValue || *limit < 1 {
		*limit = DefaultLimitValue
	}
	if *offset < DefaultOffsetValue {
		*offset = DefaultOffsetValue
	}
}

type roleUseCase struct {
	roleRepo    domain.RoleRepository
}

func NewRoleUseCase(rr domain.RoleRepository) domain.RoleRepository {
	return &roleUseCase{
		roleRepo: rr,
	}
}

func (ru roleUseCase) GetById(id model.RoleID) (res *model.Role, err error) {
	res, err = ru.roleRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ru roleUseCase) Create(name model.RoleName) (res *model.Role, err error) {
	res, err = ru.roleRepo.Create(name)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ru roleUseCase) Update(id model.RoleID, newName model.RoleName) (res *model.Role, err error) {
	res, err = ru.roleRepo.Update(id, newName)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ru roleUseCase) Delete(id model.RoleID) (err error) {
	err = ru.roleRepo.Delete(id)
	if err != nil {
		return err
	}
	return
}

func (ru roleUseCase) GetAll(limit, offset int) ([]model.Role, error) {
	checkLimitOffset(&limit, &offset)
	res, err := ru.roleRepo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	return res, nil
}