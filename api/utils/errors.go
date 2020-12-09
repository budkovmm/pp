package utils

import (
	"errors"
	"fmt"
)

const (
	userNotFoundText = "user not found"
	invalidUserIdText = "invalid user id"
	dBInternalError = "DB internal error"
    invalidRequestPayload = "invalid request payload"
    noDbInContext = "could not get database connection pool from context"
)

var UserNotFoundError = errors.New(userNotFoundText)
var InvalidUserIdError = errors.New(invalidUserIdText)
var DBInternalError = errors.New(dBInternalError)
var InvalidRequestPayload = errors.New(invalidRequestPayload)
var NoDbInContext = errors.New(noDbInContext)

func GetValidationError(msg string) error {
	errorMessage := fmt.Sprintf("Validation error: %s", msg)
	return errors.New(errorMessage)
}

