package design

import (
	. "github.com/azihsoyn/goa_microservice_sample/user_service/design/media_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("gateway", func() {
	Title("gateway api server")
	Description("sample microservice API with goa")
	Host("localhost:8081")
	Scheme("http")
})

var Complex = MediaType("application/vnd.complex+json", func() {
	Attributes(func() {
		Attribute("users", CollectionOf(User), "list of user", func() {
			Metadata("struct:field:type", "app.UserCollection", `github.com/azihsoyn/goa_microservice_sample/user_service/app`)
		})

		Required("users")
	})

	View("default", func() {
		Attribute("users")
	})
})

var _ = Resource("gateway", func() {

	BasePath("/gateway")

	Action("gateway", func() {
		Routing(
			GET(""),
		)
		Description("gateway")
		Response(OK, Complex)
	})
})
