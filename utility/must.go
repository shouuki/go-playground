package utility

import (
	"fmt"
	"runtime"
)

// errMsgFormat is the format of the error message.
const errMsgFormat = "unexpected error: %w [at %s(%s:%d)]"

// MustDo panics if err is not nil.
func MustDo(err error) {
	if err != nil {
		if pc, file, line, ok := runtime.Caller(1); !ok {
			panic(err)
		} else {
			panic(fmt.Errorf(errMsgFormat, err, runtime.FuncForPC(pc).Name(), file, line))
		}
	}
}

// MustGet panics if err is not nil or else returns the given value.
func MustGet[T any](v T, err error) T {
	if err != nil {
		if pc, file, line, ok := runtime.Caller(1); !ok {
			panic(err)
		} else {
			panic(fmt.Errorf(errMsgFormat, err, runtime.FuncForPC(pc).Name(), file, line))
		}
	}
	return v
}
