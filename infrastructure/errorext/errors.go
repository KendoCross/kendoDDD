package errorext

import (
	"encoding/json"
)

type errEmptyPayload struct{}

func (e errEmptyPayload) Error() string {
	return ""
}

type errWithCode struct {
	Code int
	Msg  string
}

func (e errWithCode) Error() string {
	return e.Msg
}

// 400 response
type BadRequestError struct {
	Name   string
	Reason string
}

func (e *BadRequestError) Error() string {
	return e.Reason
}

type BadRequestErrors []BadRequestError

func (errs BadRequestErrors) Error() (msg string) {
	for _, err := range errs {
		msg += err.Error() + "\r\n"
	}
	return
}

// 401 response
type UnauthorizedError struct {
	errEmptyPayload
}

// 403 response
type ForbiddenError struct {
	errEmptyPayload
}

// 409 response
type ConflictError struct {
	errWithCode
}

// 418 response
type TeapotError struct {
	Payload interface{}
}

func (e TeapotError) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// 500 response
type InternalServerError struct {
	errWithCode
}

func NewBadRequestError(name string, reason string) BadRequestError {
	return BadRequestError{
		Name:   name,
		Reason: reason,
	}
}

func NewInternalServerError(code int, msg string) InternalServerError {
	return InternalServerError{
		errWithCode{
			Code: code,
			Msg:  msg,
		},
	}
}

func NewConflictError(code int, msg string) ConflictError {
	return ConflictError{
		errWithCode{
			Code: code,
			Msg:  msg,
		},
	}
}

func NewTeapotError(payload interface{}) TeapotError {
	return TeapotError{
		Payload: payload,
	}
}

func NewUnauthorizedError() UnauthorizedError {
	return UnauthorizedError{}
}

func NewForbiddenError() ForbiddenError {
	return ForbiddenError{}
}