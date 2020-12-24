package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bxcodec/faker"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"pp/pkg/domain"
	"pp/pkg/domain/mocks"
	"pp/pkg/model"
	"strconv"
	"strings"
	"testing"
)

func TestGetAll(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockUCase := new(mocks.RoleUseCase)
	mockListRole := make([]model.Role, 0)
	mockListRole = append(mockListRole, mockRole)
	limit, offset := 1,1
	mockUCase.On("GetAll", limit, offset).Return(mockListRole, nil)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("limit", strconv.Itoa(limit))
	_ = writer.WriteField("offset", strconv.Itoa(offset))
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(http.MethodGet, "/role", payload)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.GetAll(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAllInternalError(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockUCase := new(mocks.RoleUseCase)
	limit, offset := 1,1
	mockUCase.On("GetAll", limit, offset).Return(nil, errors.New("some error"))
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("limit", strconv.Itoa(limit))
	_ = writer.WriteField("offset", strconv.Itoa(offset))
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(http.MethodGet, "/role", payload)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.GetAll(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("Create", mockRole.Name).Return(&mockRole, nil)
	j, err := json.Marshal(mockRole)
	assert.NoError(t, err)
	req, err := http.NewRequest(http.MethodPost, "/role", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.Create(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestCreateBadRequest(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("Create", mockRole.Name).Return(&mockRole, nil)
	req, err := http.NewRequest(http.MethodPost, "/role", strings.NewReader(``))
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.Create(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockUCase.AssertNotCalled(t, "Create")
}

func TestCreateInternalError(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("Create", mockRole.Name).Return(nil, domain.DBInternalError)
	j, err := json.Marshal(mockRole)
	assert.NoError(t, err)
	req, err := http.NewRequest(http.MethodPost, "/role", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.Create(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUCase.AssertNotCalled(t, "Create")
}

func TestGetById(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("GetById", mockRole.ID).Return(&mockRole, nil)
	payload := &bytes.Buffer{}
	req, err := http.NewRequest(http.MethodGet, "/role", payload)
	req = mux.SetURLVars(req, map[string]string{"id":"1"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.GetById(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByIdNotFoundError(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("GetById", mockRole.ID).Return(nil, domain.RoleNotFoundError)
	payload := &bytes.Buffer{}
	req, err := http.NewRequest(http.MethodGet, "/role", payload)
	req = mux.SetURLVars(req, map[string]string{"id":"1"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.GetById(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByIdInternalError(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("GetById", mockRole.ID).Return(nil, domain.DBInternalError)
	payload := &bytes.Buffer{}
	req, err := http.NewRequest(http.MethodGet, "/role", payload)
	req = mux.SetURLVars(req, map[string]string{"id":"1"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.GetById(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByIdBadRequest(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("GetById", mockRole.ID).Return(&mockRole, nil)
	payload := &bytes.Buffer{}
	req, err := http.NewRequest(http.MethodGet, "/role", payload)
	req = mux.SetURLVars(req, map[string]string{"id":"a"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.GetById(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockUCase.AssertNotCalled(t, "GetById")
}

func TestDelete(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("Delete", mockRole.ID).Return(nil)
	req, err := http.NewRequest(http.MethodPut, "/role", strings.NewReader(``))
	req = mux.SetURLVars(req, map[string]string{"id":"1"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.Delete(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestDeleteInvalidRequest(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("Delete", mockRole.ID).Return(nil)
	req, err := http.NewRequest(http.MethodPut, "/role", strings.NewReader(``))
	req = mux.SetURLVars(req, map[string]string{"id":"a"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.Delete(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockUCase.AssertNotCalled(t, "Delete")
}

func TestDeleteRoleNotFound(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("Delete", mockRole.ID).Return(domain.RoleNotFoundError)
	req, err := http.NewRequest(http.MethodPut, "/role", strings.NewReader(``))
	req = mux.SetURLVars(req, map[string]string{"id":"1"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.Delete(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestDeleteInternalError(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("Delete", mockRole.ID).Return(domain.DBInternalError)
	req, err := http.NewRequest(http.MethodPut, "/role", strings.NewReader(``))
	req = mux.SetURLVars(req, map[string]string{"id":"1"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.Delete(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("Update", mockRole.ID, mockRole.Name).Return(&mockRole, nil)
	j, err := json.Marshal(mockRole)
	assert.NoError(t, err)
	req, err := http.NewRequest(http.MethodPut, "/role", strings.NewReader(string(j)))
	req = mux.SetURLVars(req, map[string]string{"id":"1"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.Update(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestUpdateInvalidRequest(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("Update", mockRole.ID, mockRole.Name).Return(nil, domain.InvalidRequestPayload)
	req, err := http.NewRequest(http.MethodPut, "/role", strings.NewReader(``))
	req = mux.SetURLVars(req, map[string]string{"id":"1"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.Update(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockUCase.AssertNotCalled(t, "Update")
}

func TestUpdateRoleNotFound(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("Update", mockRole.ID, mockRole.Name).Return(nil, domain.RoleNotFoundError)
	j, err := json.Marshal(mockRole)
	assert.NoError(t, err)
	req, err := http.NewRequest(http.MethodPut, "/role", strings.NewReader(string(j)))
	req = mux.SetURLVars(req, map[string]string{"id":"1"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.Update(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestUpdateInternalError(t *testing.T) {
	var mockRole model.Role
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockRole.ID = 1
	mockUCase := new(mocks.RoleUseCase)
	mockUCase.On("Update", mockRole.ID, mockRole.Name).Return(nil, domain.DBInternalError)
	j, err := json.Marshal(mockRole)
	assert.NoError(t, err)
	req, err := http.NewRequest(http.MethodPut, "/role", strings.NewReader(string(j)))
	req = mux.SetURLVars(req, map[string]string{"id":"1"})
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	handler := RoleHandler{
		RoleUseCase: mockUCase,
	}
	handler.Update(w, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUCase.AssertExpectations(t)
}

func TestParseCreatePayload(t *testing.T) {
	expectedRoleName := model.RoleName("test")
	rcr := RoleRequest{Name: expectedRoleName}
	jsonValue, _ := json.Marshal(rcr)
	payload := bytes.NewBuffer(jsonValue)
	request, err := http.NewRequest(http.MethodPost, "", payload)
	assert.NoError(t, err)
	result, err := parseRolePayload(request)
	assert.NoError(t, err)
	assert.Equal(t, rcr.Name, result.Name)
}

func TestParseCreatePayloadNegative(t *testing.T) {
	payload := new(bytes.Buffer)
	request, err := http.NewRequest(http.MethodPost, "", payload)
	assert.NoError(t, err)
	_, err = parseRolePayload(request)
	assert.Error(t, err, domain.InvalidRequestPayload)
}

func TestGetLimitOffset(t *testing.T) {
	payload := new(bytes.Buffer)
	request, err := http.NewRequest(http.MethodPost, "", payload)
	assert.NoError(t, err)
	var expectedLimit, expectedOffset = 0, 0
	limit, offset := getLimitOffset(request)
	assert.Equal(t, expectedLimit, limit)
	assert.Equal(t, expectedOffset, offset)

	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("limit", "1")
	_ = writer.WriteField("offset", "1")
	err = writer.Close()
	assert.NoError(t, err)
	request, err = http.NewRequest(http.MethodGet, "", payload)
	assert.NoError(t, err)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	expectedLimit, expectedOffset = 1, 1
	limit, offset = getLimitOffset(request)
	assert.Equal(t, expectedLimit, limit)
	assert.Equal(t, expectedOffset, offset)
}