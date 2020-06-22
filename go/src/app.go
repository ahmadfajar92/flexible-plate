package src

import (
	"scaffold/config"
	"scaffold/shared/factory"
	"scaffold/shared/interfaces"
	"sync"
)

type application struct {
	cfg          *config.Config
	repositories interfaces.Repositories
	usecases     interfaces.Usecases
	deliveries   interfaces.Deliveries
}

// Application func, initialize the App
func Application(cfg *config.Config) interfaces.Application {

	app := new(application)
	app.cfg = cfg
	app.repositories = factory.Repositories(app)
	app.usecases = factory.Usecases(app)
	app.deliveries = factory.Deliveries(app)

	return app
}

func (app *application) Cfg() *config.Config {
	return app.cfg
}

func (app *application) Repository() interfaces.Repositories {
	return app.repositories
}

func (app *application) Delivery() interfaces.Deliveries {
	return app.deliveries
}

func (app *application) Usecase() interfaces.Usecases {
	return app.usecases
}

func (app *application) Run() {

	wg := new(sync.WaitGroup)

	// if stream is active then run the server
	if app.Cfg().STREAM() {
		// running stream serve
		app.ServeStream(wg)
	}

	// if rpc is active then run the server
	if app.Cfg().RPC() {
		// running rest server
		app.ServeRPC(wg)
	}

	// rest run
	if app.Cfg().REST() {
		// init echo rest server
		app.ServeRest(wg)
	}

	wg.Wait()
}
