//go:generate goagen bootstrap -d github.com/azihsoyn/goa_microservice_sample/gateway/design

package main

import (
	"github.com/azihsoyn/goa_microservice_sample/gateway/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("gateway")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "gateway" controller
	c := NewGatewayController(service)
	app.MountGatewayController(service, c)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}

}
