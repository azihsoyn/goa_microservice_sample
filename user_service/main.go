//go:generate goagen bootstrap -d github.com/azihsoyn/goa_microservice_sample/user_service/design

package main

import (
	"github.com/azihsoyn/goa_microservice_sample/user_service/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("user service")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "user" controller
	c := NewUserController(service)
	app.MountUserController(service, c)

	// Start service
	if err := service.ListenAndServe(":8082"); err != nil {
		service.LogError("startup", "err", err)
	}

}
