package utility

import (
	"fmt"
	"runtime"
)

// errMsgFormat is the format of the error message.
const errMsgFormat = "unexpected error: %s\n\tat %s(%s:%d)"

// MustDo panics if err is not nil.
func MustDo(err error) {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		if !ok {
			panic(err)
		}
		panic(fmt.Sprintf(errMsgFormat, err.Error(), runtime.FuncForPC(pc).Name(), file, line))
	}
}

// MustGet panics if err is not nil or else returns the given value.
func MustGet[T any](v T, err error) T {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		if !ok {
			panic(err)
		}
		panic(fmt.Sprintf(errMsgFormat, err.Error(), runtime.FuncForPC(pc).Name(), file, line))
	}
	return v
}
