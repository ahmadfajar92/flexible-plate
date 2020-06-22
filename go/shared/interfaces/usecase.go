package interfaces

import (
	"scaffold/shared/structure"
)

// UseCase
type UseCase interface {
	ProcessRequest(request interface{}) *structure.Result
}

// Usecases interface
type Usecases interface{}
