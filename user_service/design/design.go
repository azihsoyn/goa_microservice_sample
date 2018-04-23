package design

import (
	. "github.com/azihsoyn/goa_microservice_sample/user_service/design/media_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("user service", func() {
	Title("user api server")
	Description("sample microservice API with goa")
	Host("localhost:8082")
	Scheme("http")
})

var _ = Resource("user", func() {

	DefaultMedia(User)
	BasePath("/users")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all users.")
		Response(OK, CollectionOf(User))
	})
})
