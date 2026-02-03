package must

import (
	"log"
	"runtime"
)

const errMsgFormat = "unexpected error: %s\n\tat %s(%s:%d)"

func Do(err error) {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		if !ok {
			panic(err)
		}
		funcName := runtime.FuncForPC(pc).Name()
		log.Panicf(errMsgFormat, err.Error(), funcName, file, line)
	}
}

func Value[T any](v T, err error) T {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		if !ok {
			panic(err)
		}
		funcName := runtime.FuncForPC(pc).Name()
		log.Panicf(errMsgFormat, err.Error(), funcName, file, line)
	}
	return v
}
