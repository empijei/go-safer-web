package runtimeutils

import "runtime"

// GetCallerName tries to get the name of this function's caller caller.
// An empty string is returned in case of failure.
func GetCallerName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	details := runtime.FuncForPC(pc)
	if details == nil {
		return ""
	}
	return details.Name()

}

// GetCallerInfo returns the file, line and function name of the caller of this function caller.
// Zero data is returned in case of failure.
func GetCallerInfo() (file string, line int, fName string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return
	}
	details := runtime.FuncForPC(pc)
	if details == nil {
		return
	}
	fName = details.Name()
	return

}
