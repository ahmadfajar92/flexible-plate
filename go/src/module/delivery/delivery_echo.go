package delivery

import (
	"net/http"
	"scaffold/shared"
	"scaffold/shared/interfaces"
	"scaffold/shared/log"

	"github.com/labstack/echo"
)

type httpDelivery struct {
	app interfaces.Application
}

// GetPath func
func (d *httpDelivery) GetPath() string {
	return "/"
}

// Setup func
func Setup(application interfaces.Application) interfaces.DeliveryHTTP {
	delivery := new(httpDelivery)
	delivery.app = application

	return delivery
}

// Mount func
func (d *httpDelivery) Mount(routes interface{}) {
	group, _ := routes.(*echo.Group)

	group.GET("", d.exampleRoute)
}

func (d *httpDelivery) exampleRoute(c echo.Context) error {
	ctx := "delivery-exampleRoute"

	response := shared.JSONResponse(
		http.StatusCreated,
		"Success",
		true,
		"Yeay it's Running",
	)

	// start stream server
	log.Log(log.InfoLevel, "Delivery is Run and Response ...", ctx, "")

	return c.JSON(response.Code, response)
}
