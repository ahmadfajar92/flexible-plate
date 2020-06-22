package factory

import (
	"scaffold/shared/interfaces"
	"scaffold/shared/structure"
)

type repository struct {
	application interfaces.Application
}

// Repositories func
func Repositories(application interfaces.Application) interfaces.Repositories {
	app := new(repository)
	app.application = application
	return app
}

func (repo *repository) CommandPython() *structure.Result {
	return nil
}
