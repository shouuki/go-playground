package model

import (
	"fmt"
	"strings"
)

var (
	SystemErrorCode       = NewErrorCode("ERR0001", "{0}")
	BusinessErrorCode     = NewErrorCode("ERR0002", "{0}")
	InvalidParamErrorCode = NewErrorCode("ERR0003", "{0}")
)

type ErrorCode interface {
	Code() string
	Message(args ...any) string
}

type errorCodeImpl struct {
	code    string
	message string
}

func (e *errorCodeImpl) Code() string {
	return e.code
}

func (e *errorCodeImpl) Message(args ...any) string {
	return format(e.message, args...)
}

func format(message string, args ...any) string {
	msg := message
	for i, arg := range args {
		msg = strings.ReplaceAll(msg, fmt.Sprintf("{%d}", i), fmt.Sprintf("%v", arg))
	}
	return msg
}

func NewErrorCode(code string, message string) ErrorCode {
	return &errorCodeImpl{
		code:    code,
		message: message,
	}
}

type baseError struct {
	code  ErrorCode
	args  []any
	cause error
}

func (e *baseError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.code.Message(e.args...), e.cause)
	}
	return e.code.Message(e.args...)
}

func (e *baseError) ErrorCode() ErrorCode {
	return e.code
}

func (e *baseError) Args() []any {
	return e.args
}

func (e *baseError) Unwrap() error {
	return e.cause
}

// AppError represents all error that occurred in this application.
type AppError struct {
	baseError
}

func NewAppError(code ErrorCode, args ...any) *AppError {
	return &AppError{
		baseError: baseError{
			code:  code,
			args:  args,
			cause: nil,
		},
	}
}

func WrapAppError(cause error, code ErrorCode, args ...any) *AppError {
	return &AppError{
		baseError: baseError{
			code:  code,
			args:  args,
			cause: cause,
		},
	}
}
