package main

import (
	"encoding/json"
	"net/http"

	"github.com/azihsoyn/goa_microservice_sample/gateway/app"
	usapp "github.com/azihsoyn/goa_microservice_sample/user_service/app"
	"github.com/goadesign/goa"
)

// GatewayController implements the gateway resource.
type GatewayController struct {
	*goa.Controller
}

// NewGatewayController creates a gateway controller.
func NewGatewayController(service *goa.Service) *GatewayController {
	return &GatewayController{Controller: service.NewController("GatewayController")}
}

func FetchUsers() (usapp.UserCollection, error) {
	resp, err := http.Get("http://localhost:8082/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ret := make(usapp.UserCollection, 0)
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// Gateway runs the gateway action.
func (c *GatewayController) Gateway(ctx *app.GatewayGatewayContext) error {
	// GatewayController_Gateway: start_implement

	// Put your logic here
	us, err := FetchUsers()
	if err != nil {
		return err
	}

	res := &app.Complex{
		Users: us,
	}
	return ctx.OK(res)
	// GatewayController_Gateway: end_implement
}
