package utility

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// CurrentRoutineId return current goroutine id, for example 1.
func CurrentRoutineId() int64 {
	id := 0
	buf := make([]byte, 32)
	runtime.Stack(buf, false)
	buf = buf[:bytes.LastIndexByte(buf, ' ')]
	_, _ = fmt.Sscanf(string(buf), "goroutine %d", &id)
	return int64(id)
}

// CurrentRoutineName return current goroutine name, for example goroutine 1.
func CurrentRoutineName() string {
	return fmt.Sprintf("goroutine %d", CurrentRoutineId())
}

// CurrentPackageName return current package import path, for example go-playground/utility.
func CurrentPackageName() (string, error) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("package name not resolvable")
	}
	modulePath := ""
	packageDir := ""
	funcName := runtime.FuncForPC(pc).Name()
	slashIndex := strings.LastIndex(funcName, "/")
	if slashIndex != -1 {
		modulePath = funcName[:slashIndex]
		funcName = funcName[slashIndex:]
	}
	packageDir, _, _ = strings.Cut(funcName, ".")
	return modulePath + packageDir, nil
}
