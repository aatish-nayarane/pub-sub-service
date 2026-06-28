package errorHandler

import (
	"fmt"
	"path"
	"runtime"
)

type Category string

const (
	// API services error categories
	CategoryAPI Category = "API"

	// Corelibs error categories
	CategoryDB            Category = "DB"
	CategoryENV           Category = "ENV"
	CategoryBackgroundJob Category = "BACKGROUND_JOB"
	CategoryAPICall       Category = "API_CALL"
	LibPublic             Category = "LIB_PUBLIC"
	LibPrivate            Category = "LIB_PRIVATE"
	CategoryQuotaExceeded Category = "QUOTA_EXCEEDED"
)

type Error struct {
	C       Category
	Op      string
	Message string
	Err     error
	File    string
	Line    int
}

func (e *Error) Error() string {
	base := fmt.Sprintf("[%s][%s] %s (%s:%d)", e.C, e.Op, e.Message, path.Base(e.File), e.Line)
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", base, e.Err)
	}
	return base
}

func (e *Error) Unwrap() error {
	return e.Err
}

// New creates a new *Error capturing file/line
func New(c Category, op, msg string, err error, file string, line int) *Error {
	return &Error{C: c, Op: op, Message: msg, Err: err, File: file, Line: line}
}

// Wrap wraps err with new context but preserves original Error if present
func Wrap(c Category, op, msg string, err error) *Error {
	_, file, line, _ := runtime.Caller(1)
	if err == nil {
		return nil
	}
	return New(c, op, msg, err, file, line)
}

// // Example usage in your code:
// func UpdateSomething() error {
// 	const op = "pkg.Change.UpdateSomething"
// 	rawErr := doDBOp()
// 	if rawErr != nil {
// 		return Wrap(CategoryDB, op, "updating record failed", rawErr)
// 	}
// 	return nil
// }
