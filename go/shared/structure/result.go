package structure

import "errors"

// Result struct
type Result struct {
	Data    interface{}
	Error   error
	Message string
}

// ErrorResult func return an error result
func ErrorResult(msg string) *Result {
	return &Result{
		Error: errors.New(msg),
	}
}
