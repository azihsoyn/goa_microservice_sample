package media_types

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var User = MediaType("application/vnd.user+json", func() {
	Description("user")
	Attributes(func() {
		Attribute("id", Integer, "ID of user", func() {
			Example(1)
		})
		Attribute("name", String, "Name of user", func() {
			Example("test")
		})
		Attribute("created_at", DateTime, "Date of creation")

		Required("id", "name", "created_at")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("created_at")
	})
})
