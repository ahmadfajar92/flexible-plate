package src

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (app *application) ServeRest(wg *sync.WaitGroup) {
	app.runEchoServer(wg)
}

// EchoRESTServer func
func (app *application) runEchoServer(wg *sync.WaitGroup) {
	e := echo.New()
	e.Debug = app.Cfg().Debug()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(sharedMiddleware.Authorization)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	deliveries := app.deliveries.GetRest()
	for _, delivery := range deliveries {
		routes := e.Group(delivery.GetPath())
		// just mount the route
		delivery.Mount(routes)
	}

	wg.Add(1)
	port := fmt.Sprintf(":%d", app.Cfg().Port())
	e.Logger.Fatal(e.Start(port))
	wg.Done()

	return
}
