package usecase

import (
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pp/pkg/domain"
	"pp/pkg/domain/mocks"
	"pp/pkg/model"
	"testing"
)

func TestGetAll(t *testing.T) {
	mockRoleRepo := new(mocks.RoleRepository)
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRoleList := make([]model.Role, 0)
	mockRoleList = append(mockRoleList, mockRole)
	
	t.Run("successful", func(t *testing.T) {
		mockRoleRepo.On("GetAll",  mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockRoleList, nil).Once()
		u := NewRoleUseCase(mockRoleRepo)
		limit, offset := 1,1
		roleList, err := u.GetAll(limit, offset)
		assert.NoError(t, err)
		assert.Len(t, roleList, len(mockRoleList))
		mockRoleRepo.AssertExpectations(t)
	})

	t.Run("negative", func(t *testing.T) {
		mockRoleRepo.On("GetAll",  mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, domain.DBInternalError).Once()
		u := NewRoleUseCase(mockRoleRepo)
		limit, offset := 1,1
		_, err := u.GetAll(limit, offset)
		assert.Error(t, err, domain.DBInternalError)
		mockRoleRepo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mockRoleRepo := new(mocks.RoleRepository)
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)

	t.Run("successful", func(t *testing.T) {
		mockRoleRepo.On("GetById",  mock.AnythingOfType("model.RoleID")).Return(&mockRole, nil).Once()
		u := NewRoleUseCase(mockRoleRepo)
		role, err := u.GetById(mockRole.ID)
		assert.NoError(t, err)
		assert.Equal(t, &mockRole, role)
		mockRoleRepo.AssertExpectations(t)
	})

	t.Run("negative", func(t *testing.T) {
		mockRoleRepo.On("GetById",  mock.AnythingOfType("model.RoleID")).Return(nil, domain.DBInternalError).Once()
		u := NewRoleUseCase(mockRoleRepo)
		_, err := u.GetById(mockRole.ID)
		assert.Error(t, err, domain.DBInternalError)
		mockRoleRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockRoleRepo := new(mocks.RoleRepository)
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)

	t.Run("successful", func(t *testing.T) {
		mockRoleRepo.On("Create",  mock.AnythingOfType("model.RoleName")).Return(&mockRole, nil).Once()
		u := NewRoleUseCase(mockRoleRepo)
		role, err := u.Create(mockRole.Name)
		assert.NoError(t, err)
		assert.Equal(t, &mockRole, role)
		mockRoleRepo.AssertExpectations(t)
	})

	t.Run("negative", func(t *testing.T) {
		mockRoleRepo.On("Create",  mock.AnythingOfType("model.RoleName")).Return(nil, domain.DBInternalError).Once()
		u := NewRoleUseCase(mockRoleRepo)
		_, err := u.Create(mockRole.Name)
		assert.Error(t, err, domain.DBInternalError)
		mockRoleRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockRoleRepo := new(mocks.RoleRepository)
	newRoleName := model.RoleName("test_role_name")
	var mockRole model.Role
	mockRole.Name = newRoleName
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)

	t.Run("successful", func(t *testing.T) {
		mockRoleRepo.On("Update", mock.AnythingOfType("model.RoleID"), mock.AnythingOfType("model.RoleName")).Return(&mockRole, nil).Once()
		u := NewRoleUseCase(mockRoleRepo)
		role, err := u.Update(mockRole.ID, newRoleName)
		assert.NoError(t, err)
		assert.Equal(t, mockRole.Name, role.Name)
		mockRoleRepo.AssertExpectations(t)
	})

	t.Run("negative", func(t *testing.T) {
		mockRoleRepo.On("Update", mock.AnythingOfType("model.RoleID"), mock.AnythingOfType("model.RoleName")).Return(nil, domain.DBInternalError).Once()
		u := NewRoleUseCase(mockRoleRepo)
		_, err := u.Update(mockRole.ID, newRoleName)
		assert.Error(t, err, domain.DBInternalError)
		mockRoleRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockRoleRepo := new(mocks.RoleRepository)
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)

	t.Run("successful", func(t *testing.T) {
		mockRoleRepo.On("Delete",  mock.AnythingOfType("model.RoleID")).Return(nil).Once()
		u := NewRoleUseCase(mockRoleRepo)
		err := u.Delete(mockRole.ID)
		assert.NoError(t, err)
		mockRoleRepo.AssertExpectations(t)
	})

	t.Run("negative", func(t *testing.T) {
		mockRoleRepo.On("Delete",  mock.AnythingOfType("model.RoleID")).Return(domain.DBInternalError).Once()
		u := NewRoleUseCase(mockRoleRepo)
		err := u.Delete(mockRole.ID)
		assert.Error(t, err, domain.DBInternalError)
		mockRoleRepo.AssertExpectations(t)
	})
}

func TestCheckLimitOffset(t *testing.T) {
	var limit, offset = 0, 0
	var expectedLimit, expectedOffset = 10, 0
	checkLimitOffset(&limit, &offset)
	assert.Equal(t, expectedLimit, limit)
	assert.Equal(t, expectedOffset, offset)

	limit, offset = 10, 0
	checkLimitOffset(&limit, &offset)
	assert.Equal(t, expectedLimit, limit)
	assert.Equal(t, expectedOffset, offset)

	limit, offset = 0, -1
	checkLimitOffset(&limit, &offset)
	assert.Equal(t, expectedLimit, limit)
	assert.Equal(t, expectedOffset, offset)

	limit, offset = 1, 0
	expectedLimit, expectedOffset = 1, 0
	checkLimitOffset(&limit, &offset)
	assert.Equal(t, expectedLimit, limit)
	assert.Equal(t, expectedOffset, offset)

	limit, offset = 11, 0
	expectedLimit, expectedOffset = 10, 0
	checkLimitOffset(&limit, &offset)
	assert.Equal(t, expectedLimit, limit)
	assert.Equal(t, expectedOffset, offset)
}
