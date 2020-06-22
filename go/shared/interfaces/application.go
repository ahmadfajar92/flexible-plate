package interfaces

import (
	"scaffold/config"
	"sync"
)

// Application interface
type Application interface {
	// config
	Cfg() *config.Config
	// factory
	// Delivery() Deliveries
	Repository() Repositories
	// serve request
	ServeRest(wg *sync.WaitGroup)
	ServeStream(wg *sync.WaitGroup)
	ServeRPC(wg *sync.WaitGroup)
	// run
	Run()
}
