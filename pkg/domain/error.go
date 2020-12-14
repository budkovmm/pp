package domain

import (
	"errors"
	"fmt"
)

const (
	roleNotFoundText      = "role not found"
	invalidUserIdText     = "invalid user id"
	dBInternalError       = "DB internal error"
	invalidRequestPayload = "invalid request payload"
)

var RoleNotFoundError = errors.New(roleNotFoundText)
var InvalidUserIdError = errors.New(invalidUserIdText)
var DBInternalError = errors.New(dBInternalError)
var InvalidRequestPayload = errors.New(invalidRequestPayload)

func GetValidationError(msg string) error {
	errorMessage := fmt.Sprintf("Validation error: %s", msg)
	return errors.New(errorMessage)
}
