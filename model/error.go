package model

import (
	"fmt"
	"strings"
)

var (
	SystemError       = NewErrorCode("ERR0001", "System Error: {0}")
	InvalidParamError = NewErrorCode("ERR0002", "Invalid Parameter: {0}")
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
