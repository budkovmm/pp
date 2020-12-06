package utils

import "errors"

const (
	userNotFoundText = "user not found"
	invalidUserIdText = "invalid user id"
	dBInternalError = "DB internal error"
    invalidRequestPayload = "invalid request payload"
)

var UserNotFoundError = errors.New(userNotFoundText)
var InvalidUserIdError = errors.New(invalidUserIdText)
var DBInternalError = errors.New(dBInternalError)
var InvalidRequestPayload = errors.New(invalidRequestPayload)

