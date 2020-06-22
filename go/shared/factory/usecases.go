package factory

import "scaffold/shared/interfaces"

type usecases struct {
	application interfaces.Application
}

// Usecases func
func Usecases(application interfaces.Application) interfaces.Usecases {
	ucs := new(usecases)
	ucs.application = application
	return ucs
}
