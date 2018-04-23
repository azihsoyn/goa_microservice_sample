package main

import (
	"time"

	"github.com/azihsoyn/goa_microservice_sample/user_service/app"
	"github.com/goadesign/goa"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// UserController_List: start_implement

	// Put your logic here

	res := app.UserCollection{
		{ID: 1, Name: "Andrew", CreatedAt: time.Now()},
		{ID: 2, Name: "Bob", CreatedAt: time.Now()},
		{ID: 3, Name: "Charlse", CreatedAt: time.Now()},
	}
	return ctx.OK(res)
	// UserController_List: end_implement
}
